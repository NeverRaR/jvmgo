package constants

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (receiver *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
