package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
func (self *Method) copyMemberInfo(cfMethod *classfile.MemberInfo) {
	self.ClassMember.class = self.class
	self.ClassMember.name = cfMethod.Name()
	self.ClassMember.descriptor = cfMethod.Descriptor()
	self.ClassMember.accessFlags = cfMethod.AccessFlags()
}

//getters
func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) IsPublic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_PUBLIC
}

func (self *Method) IsFinal() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_FINAL
}

func (self *Method) IsAbstract() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_ABSTRACT
}

func (self *Method) IsSynthetic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_SYNTHETIC
}

func (self *Method) IsStatic() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_STATIC
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) Descriptor() string {
	return self.descriptor
}
