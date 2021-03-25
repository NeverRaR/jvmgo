package heap

func (receiver *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, receiver

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}
}

// receiver extends c
func (receiver *Class) IsSubClassOf(other *Class) bool {
	for c := receiver.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// receiver implements iface
func (receiver *Class) IsImplements(iface *Class) bool {
	for c := receiver; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// receiver extends iface
func (receiver *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range receiver.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends receiver
func (receiver *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(receiver)
}

// iface extends receiver
func (receiver *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.IsSubInterfaceOf(receiver)
}
