package comparisons

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IFEQ struct {
	base.BranchInstruction
}

func (receiver *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, receiver.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (receiver *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, receiver.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (receiver *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, receiver.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (receiver *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, receiver.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (receiver *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, receiver.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (receiver *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, receiver.Offset)
	}
}
