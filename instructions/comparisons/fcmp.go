package comparisons

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type FCMPG struct {
	base.NoOperandsInstruction
}

func (receiver *FCMPG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

type FCMPL struct {
	base.NoOperandsInstruction
}

func (receiver *FCMPL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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
