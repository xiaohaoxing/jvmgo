package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}
//把传入的classData解析成ClassFilestruct,使用了下面的read方法
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
//依次调用下面的各个read方法，按照：
//Magic,version,constantPool,accessFlags,
//thisClass,superClass,interfaces,
//fields,methods,attributes的顺序读取。
//说明Class文件里面的信息按照这个顺序排列的。
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}
//class文件前8个byte为：0xCAFEBABE作为标识
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}
//magic之后是u2类型的次版本和主版本
//这里假设我们跟Java8一样，支持45.0~52.0版本的class文件。
//其中45.X有多个，其他主版本的次版本只有0
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50,51,52:
		if self.minorVersion == 0{
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
//版本号之后是常量池。比较复杂
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
//常量池之后是类访问标识(AccessFlags)，表示这个class是类还是接口，是public还是protected。这里暂时不验证。
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
//接口索引表后面是字段表和方法表。
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}
//类访问标识之后是2个u2类型的常量池索引，分别是thisClass和superClass

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
//只有object的superClass的常量池索引是0
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""//只有Object会运行到这里，他没有superClass
}
//在类和超类索引后面是接口索引表(Interfaces)
//从常量池查询接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}