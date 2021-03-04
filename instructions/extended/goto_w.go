package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type GOTO_W struct {
	offset int
}

func (receiver *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	receiver.offset = int(reader.ReadInt32())
}
func (receiver *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, receiver.offset)
}
