package heap

import "jvmgo/ch09/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		//methods[i] = &Method{}
		//methods[i].class = class
		//methods[i].copyMemberInfo(cfMethod)
		//methods[i].copyAttributes(cfMethod)
		//methods[i].calcArgSlotCount()
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method

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

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} //return
	case 'D':
		self.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
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

func (self *Method) IsNative() bool {
	return 0 != self.ClassMember.AccessFlags()&ACC_NATIVE
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) Descriptor() string {
	return self.descriptor
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}

}
