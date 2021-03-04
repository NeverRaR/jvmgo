package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (receiver *IINC) FetchOperands(reader *base.BytecodeReader) {
	receiver.Index = uint(reader.ReadUint8())
	receiver.Const = int32(reader.ReadInt8())
}

func (receiver *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(receiver.Index)
	val += receiver.Const
	localVars.SetInt(receiver.Index, val)
}
