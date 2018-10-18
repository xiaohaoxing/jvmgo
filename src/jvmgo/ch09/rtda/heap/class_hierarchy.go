package heap

func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

func (self *Class) IsSubInterfaceOf(other *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == other || superInterface.IsSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperInterfaceOf(other *Class) bool {
	return other.IsSubInterfaceOf(self)
}
func (self *Class) IsImplements(iface *Class) bool {
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
	if !s.IsArray() {
		// s不是数组
		if !s.IsInterface() {
			// s 不是接口类
			if !t.IsInterface() {
				// t 也不是接口类，判断 s 是否是 t 的子类
				return s.IsSubClassOf(t)
			} else {
				// t 是接口类，判断 s 是否是 t 的实现
				return s.IsImplements(t)
			}
		} else {
			// s 是接口类
			if !t.IsInterface() {
				// t 不是接口类
				return t.isJlObject()
			} else {
				// t 也是接口类，判断 t 是否是 s 的超接口
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else {
		// s 是数组
		if !t.IsArray() {
			// t 不是数组
			if !t.IsInterface() {
				// t 不是接口
				return t.isJlObject()
			} else {
				// t 是接口
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t 是数组,
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
	return false
}
