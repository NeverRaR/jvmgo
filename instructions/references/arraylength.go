package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (receiver *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
