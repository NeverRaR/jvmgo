package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (receiver *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, receiver.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (receiver *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, receiver.Offset)
	}
}
