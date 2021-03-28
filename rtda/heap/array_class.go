package heap

func (receiver *Class) IsArray() bool {
	return receiver.name[0] == '['
}

func (receiver *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(receiver.name)
	return receiver.loader.LoadClass(componentClassName)
}

func (receiver *Class) NewArray(count uint) *Object {
	if !receiver.IsArray() {
		panic("Not array class: " + receiver.name)
	}
	switch receiver.Name() {
	case "[Z":
		return &Object{receiver, make([]int8, count), nil}
	case "[B":
		return &Object{receiver, make([]int8, count), nil}
	case "[C":
		return &Object{receiver, make([]uint16, count), nil}
	case "[S":
		return &Object{receiver, make([]int16, count), nil}
	case "[I":
		return &Object{receiver, make([]int32, count), nil}
	case "[J":
		return &Object{receiver, make([]int64, count), nil}
	case "[F":
		return &Object{receiver, make([]float32, count), nil}
	case "[D":
		return &Object{receiver, make([]float64, count), nil}
	default:
		return &Object{receiver, make([]*Object, count), nil}
	}
}

func NewByteArray(loader *ClassLoader, bytes []int8) *Object {
	return &Object{loader.LoadClass("[B"), bytes, nil}
}
