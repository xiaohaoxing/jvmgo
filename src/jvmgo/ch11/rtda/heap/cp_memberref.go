package heap

import "jvmgo/ch11/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

//
//func (self *MemberRef) copyAttributes(cfField *classfile.MemberInfo) {
//	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
//		self.constValueIndex = uint(valAttr.ConstantValueIndex())
//	}
//}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
