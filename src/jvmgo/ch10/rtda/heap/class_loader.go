package heap

import (
	"fmt"
	"jvmgo/ch10/classfile"
	"jvmgo/ch10/classpath"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class //加载了的类
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
	loader.LoadBasicClasses()
	loader.LoadPrimitiveClasses()
	return loader
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class //已经加载了
	}
	var class *Class
	if name[0] == '[' {
		class = self.loadArrayClass(name)
	} else {
		class = self.loadNonArrayClass(name)
	}
	if jlClassClass, ok := self.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

//加载一个数组类
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // TODO
		name:        name,
		loader:      self,
		initStarted: true,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	//1. 找到 class 文件并把数据读取到内存
	data, entry := self.readClass(name)
	//2. 解析 class 文件，生成类数据并放入方法区
	class := self.defineClass(data)
	//3. 链接
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

//1. 读取 class 文件
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

//2. 解析 class 文件
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class) //加载父类
	resolveInterfaces(class) //加载接口
	self.classMap[class.name] = class
	return class
}
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

//递归加载父类和接口
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

//3. 链接类
func link(class *Class) {
	verify(class)  //验证
	prepare(class) //准备
}

func verify(class *Class) {
	//为了确保安全性，这里要对 java 类严格的验证。这里略过了。
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//计算对象的个数，并给他们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//是否是 long 或 double，以便占2个位置
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

//计算静态成员的个数并给他们编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I": //占一格的,byte,char,short,int
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J": //long 类型
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F": //float 类型
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D": //double 类型
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String": //字符串常量
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}

func (self *ClassLoader) LoadBasicClasses() {
	jlClassClass := self.LoadClass("java/lang/Class")
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (self *ClassLoader) LoadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		// 加载 void， int， float 等
		self.loadPrimitiveClass(primitiveType)
	}
}

// 加载基本类型，基本类型没有超类也没有接口
func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name:        className,
		loader:      self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
}
