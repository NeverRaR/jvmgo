package stack

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type POP struct {
	base.NoOperandsInstruction
}

func (receiver *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (receiver *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
