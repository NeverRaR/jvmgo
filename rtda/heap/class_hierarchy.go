package heap

func (receiver *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, receiver

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
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
