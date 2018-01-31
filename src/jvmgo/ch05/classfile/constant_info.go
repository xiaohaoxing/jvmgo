package classfile
import "fmt"
//Java虚拟机规定了14种常量
const (
	CONSTANT_Utf8 = 1
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType = 12
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)
/*
定义一个常量池中的信息，结构如下：
cp-info {
	u1 tag;		//区分常量类型字段
	u1 info[];
}
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}
//常量池中信息的构造函数，根据tag的不同类型创建不同类型的ConstantInfo
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch(tag) {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return &ConstantFloatInfo{}
	case CONSTANT_Long: return &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}
	case CONSTANT_Utf8: return &ConstantUtf8Info{}
	case CONSTANT_String: return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref: return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref: return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref: return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
	default:panic(fmt.Errorf("java.lang.ClassFormatError: constant pool tag:%v!", tag))
	}
}