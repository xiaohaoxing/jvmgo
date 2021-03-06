package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) Code() []byte {
	return self.method.Code()
}

func (self *MethodRef) Descriptor() string {
	return self.method.Descriptor()
}

func (self *MethodRef) Name() string {
	return self.method.Name()
}
