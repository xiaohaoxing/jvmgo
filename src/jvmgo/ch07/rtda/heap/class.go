package heap

import (
	"jvmgo/ch07/classfile"
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
	return class
}

func (self *Class) Name() string {
	return self.name
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

/*
实例化一个Object
*/
func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, desp string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == desp {
			return method
		}
	}
	return nil
}
