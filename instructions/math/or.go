package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (receiver *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (receiver *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
