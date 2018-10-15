package heap

import "jvmgo/ch07/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

//实例化一个成员变量
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (self *Field) IsPublic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_PUBLIC
}

func (self *Field) IsFinal() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_FINAL
}

func (self *Field) IsSynthetic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_SYNTHETIC
}

func (self *Field) IsEnum() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_ENUM
}

func (self *Field) IsStatic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_STATIC
}

//getter
func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
