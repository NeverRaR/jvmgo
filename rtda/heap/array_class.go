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
		return &Object{receiver, make([]int8, count)}
	case "[B":
		return &Object{receiver, make([]int8, count)}
	case "[C":
		return &Object{receiver, make([]uint16, count)}
	case "[S":
		return &Object{receiver, make([]int16, count)}
	case "[I":
		return &Object{receiver, make([]int32, count)}
	case "[J":
		return &Object{receiver, make([]int64, count)}
	case "[F":
		return &Object{receiver, make([]float32, count)}
	case "[D":
		return &Object{receiver, make([]float64, count)}
	default:
		return &Object{receiver, make([]*Object, count)}
	}
}
