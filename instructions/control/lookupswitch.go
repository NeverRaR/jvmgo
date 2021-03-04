package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (receiver *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	receiver.defaultOffset = reader.ReadInt32()
	receiver.matchOffsets = reader.ReadInt32s(receiver.npairs * 2)
}

//todo improve performance
func (receiver *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < receiver.npairs*2; i += 2 {
		if receiver.matchOffsets[i] == key {
			offset := receiver.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(receiver.defaultOffset))
}
