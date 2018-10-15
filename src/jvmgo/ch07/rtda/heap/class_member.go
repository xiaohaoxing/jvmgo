package heap

import "jvmgo/ch07/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

//从 class 信息中得到 member 的信息
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
