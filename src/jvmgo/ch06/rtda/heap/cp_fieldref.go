package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

/*
如何通过一个字段引用去搜索一个成员变量
*/
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError!")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError!")
	}
	self.field = field
}

/*
从一个类中查找指定变量名和描述的字段，需要分成三步找
1. 找本类的各个 members
2. 递归找 interface 的各个 members
3. 递归找父类的各个 memebers
*/
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
