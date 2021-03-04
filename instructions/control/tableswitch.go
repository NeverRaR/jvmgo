package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (receiver *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	receiver.defaultOffset = reader.ReadInt32()
	receiver.low = reader.ReadInt32()
	receiver.high = reader.ReadInt32()
	jumpOffsetsCount := receiver.high - receiver.low + 1
	receiver.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (receiver *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= receiver.low && index <= receiver.high {
		offset = int(receiver.jumpOffsets[index-receiver.low])
	} else {
		offset = int(receiver.defaultOffset)
	}
	base.Branch(frame, offset)
}
