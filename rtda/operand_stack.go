package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (receiver *OperandStack) PushInt(val int32) {
	receiver.slots[receiver.size].num = val
	receiver.size++
}
func (receiver *OperandStack) PopInt() int32 {
	receiver.size--
	return receiver.slots[receiver.size].num
}
func (receiver *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	receiver.slots[receiver.size].num = int32(bits)
	receiver.size++
}
func (receiver *OperandStack) PopFloat() float32 {
	receiver.size--
	bits := uint32(receiver.slots[receiver.size].num)
	return math.Float32frombits(bits)
}
func (receiver *OperandStack) PushLong(val int64) {
	receiver.slots[receiver.size].num = int32(val)
	receiver.slots[receiver.size+1].num = int32(val >> 32)
	receiver.size += 2
}
func (receiver *OperandStack) PopLong() int64 {
	receiver.size -= 2
	low := uint32(receiver.slots[receiver.size].num)
	high := uint32(receiver.slots[receiver.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (receiver *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	receiver.PushLong(int64(bits))
}
func (receiver *OperandStack) PopDouble() float64 {
	bits := uint64(receiver.PopLong())
	return math.Float64frombits(bits)
}
func (receiver *OperandStack) PushRef(ref *Object) {
	receiver.slots[receiver.size].ref = ref
	receiver.size++
}
func (receiver *OperandStack) PopRef() *Object {
	receiver.size--
	ref := receiver.slots[receiver.size].ref
	receiver.slots[receiver.size].ref = nil
	return ref
}
func (receiver *OperandStack) PushSlot(slot Slot) {
	receiver.slots[receiver.size] = slot
	receiver.size++
}
func (receiver *OperandStack) PopSlot() Slot {
	receiver.size--
	return receiver.slots[receiver.size]
}
