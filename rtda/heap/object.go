package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func (receiver *Object) Data() interface{} {
	return receiver.data
}

func (receiver *Object) Extra() interface{} {
	return receiver.extra
}

func (receiver *Object) Class() *Class {
	return receiver.class
}

func (receiver *Object) Fields() Slots {
	return receiver.data.(Slots)
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (receiver *Object) GetRefVar(name, descriptor string) *Object {
	field := receiver.class.getField(name, descriptor, false)
	slots := receiver.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (receiver *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(receiver.class)
}

func (receiver *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := receiver.class.getField(name, descriptor, false)
	slots := receiver.Fields()
	slots.SetRef(field.slotId, ref)
}
