package heap

import (
	"jvmgo/ch11/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string //this class' name
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object //指向 Java 类对象
	sourceFile        string
}

/*
把  ClassFile 对象转化为 Class 对象，
对其中的 constantPool、field、methods 都进行进一步读取。
*/
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) Fields() []*Field {
	return self.fields
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

/*
可访问性，要么是 public，要么是同一个包的类。（类不存在 private，所有类最低是包访问性）
*/
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == self.GetPackageName()
}

/*
从全名中获取包名，如果 package 缺失，默认包就是空字符串
如类全名是：java/lang/Object，包就是java/lang
*/
func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) GetJClass() *Object {
	return self.jClass
}

/*
实例化一个Object
*/
func (self *Class) NewObject() *Object {
	return newObject(self)
}

// 获取程序入口函数
func (self *Class) GetMainMethod() *Method {
	return self.getMethod("main", "([Ljava/lang/String;)V", true)
}

// 获取类初始化方法
func (self *Class) GetClinitMethod() *Method {
	return self.getMethod("<clinit>", "()V", true)
}

// 获取指定函数名称和参数，返回值的方法
// 包括 isStatic 的 Flag
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

// 从一个对象开始向父类向上找，直到找到给定的成员变量
func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.Descriptor() == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetInstanceMethod(methodName, methodDescriptor string) *Method {
	return self.getMethod(methodName, methodDescriptor, false)
}

func (self *Class) GetStaticMethod(methodName, methodDescriptor string) *Method {
	return self.getMethod(methodName, methodDescriptor, true)
}
func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}

func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}
