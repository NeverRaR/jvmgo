package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (receiver Slots) SetInt(index uint, val int32) {
	receiver[index].num = val
}
func (receiver Slots) GetInt(index uint) int32 {
	return receiver[index].num
}

func (receiver Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	receiver[index].num = int32(bits)
}
func (receiver Slots) GetFloat(index uint) float32 {
	bits := uint32(receiver[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (receiver Slots) SetLong(index uint, val int64) {
	receiver[index].num = int32(val)
	receiver[index+1].num = int32(val >> 32)
}
func (receiver Slots) GetLong(index uint) int64 {
	low := uint32(receiver[index].num)
	high := uint32(receiver[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (receiver Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	receiver.SetLong(index, int64(bits))
}
func (receiver Slots) GetDouble(index uint) float64 {
	bits := uint64(receiver.GetLong(index))
	return math.Float64frombits(bits)
}

func (receiver Slots) SetRef(index uint, ref *Object) {
	receiver[index].ref = ref
}
func (receiver Slots) GetRef(index uint) *Object {
	return receiver[index].ref
}
