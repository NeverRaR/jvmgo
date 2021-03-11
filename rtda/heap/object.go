package heap

type Object struct {
	class  *Class
	fields Slots
}

func (receiver *Object) Class() *Class {
	return receiver.class
}

func (receiver *Object) Fields() Slots {
	return receiver.fields
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (receiver *Object) IsInstanceOf(class *Class) bool {
	return class.isAccessibleTo(receiver.class)
}
