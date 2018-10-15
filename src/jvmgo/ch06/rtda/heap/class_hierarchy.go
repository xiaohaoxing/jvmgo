package heap

func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInstanceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInstanceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInstanceOf(iface) {
			return true
		}
	}
	return false
}

// 是否可以类型转换。如果 T 是接口。看 S 是否实现该接口。如果 T 是类，则看 S 是否是其子类。
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}
