package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (receiver *DCMPG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(1)
	}
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (receiver *DCMPL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(-1)
	}
}
