package loads

import "jvmgo/instructions/base"
import "jvmgo/rtda"

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct {
	base.Index8Instruction
}

func (receiver *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(receiver.Index))
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (receiver *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (receiver *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (receiver *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (receiver *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
