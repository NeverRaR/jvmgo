package heap

func (receiver *Object) Bytes() []int8 {
	return receiver.data.([]int8)
}
func (receiver *Object) Shorts() []int16 {
	return receiver.data.([]int16)
}
func (receiver *Object) Ints() []int32 {
	return receiver.data.([]int32)
}
func (receiver *Object) Longs() []int64 {
	return receiver.data.([]int64)
}
func (receiver *Object) Chars() []uint16 {
	return receiver.data.([]uint16)
}
func (receiver *Object) Floats() []float32 {
	return receiver.data.([]float32)
}
func (receiver *Object) Doubles() []float64 {
	return receiver.data.([]float64)
}
func (receiver *Object) Refs() []*Object {
	return receiver.data.([]*Object)
}

func (receiver *Object) ArrayLength() int32 {
	switch receiver.data.(type) {
	case []int8:
		return int32(len(receiver.data.([]int8)))
	case []int16:
		return int32(len(receiver.data.([]int16)))
	case []int32:
		return int32(len(receiver.data.([]int32)))
	case []int64:
		return int32(len(receiver.data.([]int64)))
	case []uint16:
		return int32(len(receiver.data.([]uint16)))
	case []float32:
		return int32(len(receiver.data.([]float32)))
	case []float64:
		return int32(len(receiver.data.([]float64)))
	case []*Object:
		return int32(len(receiver.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
