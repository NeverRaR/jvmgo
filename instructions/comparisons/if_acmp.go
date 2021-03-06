package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (receiver *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, receiver.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (receiver *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, receiver.Offset)
	}
}
