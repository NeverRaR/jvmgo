package rtda

import (
	"jvmgo/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}
func (receiver LocalVars) SetInt(index uint, val int32) {
	receiver[index].num = val
}
func (receiver LocalVars) GetInt(index uint) int32 {
	return receiver[index].num
}
func (receiver LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	receiver[index].num = int32(bits)
}
func (receiver LocalVars) GetFloat(index uint) float32 {
	bits := uint32(receiver[index].num)
	return math.Float32frombits(bits)
}
func (receiver LocalVars) SetLong(index uint, val int64) {
	receiver[index].num = int32(val)
	receiver[index+1].num = int32(val >> 32)
}
func (receiver LocalVars) GetLong(index uint) int64 {
	low := uint32(receiver[index].num)
	high := uint32(receiver[index+1].num)
	return int64(high)<<32 | int64(low)
}
func (receiver LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	receiver.SetLong(index, int64(bits))
}
func (receiver LocalVars) GetDouble(index uint) float64 {
	bits := uint64(receiver.GetLong(index))
	return math.Float64frombits(bits)
}
func (receiver LocalVars) SetRef(index uint, ref *heap.Object) {
	receiver[index].ref = ref
}
func (receiver LocalVars) GetRef(index uint) *heap.Object {
	return receiver[index].ref
}
