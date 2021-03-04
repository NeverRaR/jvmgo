package constants

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (receiver *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	receiver.val = reader.ReadInt8()
}

func (receiver *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(receiver.val)
	frame.OperandStack().PushInt(i)
}

func (receiver *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	receiver.val = reader.ReadInt16()
}

func (receiver *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(receiver.val)
	frame.OperandStack().PushInt(i)
}
