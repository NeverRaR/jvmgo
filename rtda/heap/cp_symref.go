package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (receiver *SymRef) ResolvedClass() *Class {
	if receiver.class == nil {
		receiver.resolveClassRef()
	}
	return receiver.class
}

func (receiver *SymRef) resolveClassRef() {
	d := receiver.cp.class
	c := d.loader.LoadClass(receiver.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	receiver.class = c
}
