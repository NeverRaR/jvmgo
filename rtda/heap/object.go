package heap

type Object struct {
	class *Class
	data  interface{}
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

func (receiver *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(receiver.class)
}
