package classfile
/*
常量池的解析
注意：
1. 第一个位置给的常量池的大小比实际大小多1
2. 索引从1~n-1，0是无效的
3. CONSTANT_Long_info,CONSTANT_Double_info各占2个位置，比较特殊
*/
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i ++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo://这2个常量信息占2个位置
			i++
		}
	}
}
//从常量池中返回指定index的一个
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
//获取指定index的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
//常量池中找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
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
	case CONSTANT_Methodref: return &ConstantMethodrefInfo{ConstatnMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref: return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
	default:panic("java.lang.ClassFormatError: constant pool tag!")
	}
}