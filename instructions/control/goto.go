package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type GOTO struct {
	base.BranchInstruction
}

func (receiver *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, receiver.Offset)
}
