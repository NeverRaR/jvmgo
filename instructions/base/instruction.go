package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (receiver *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (receiver *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	receiver.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (receiver *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	receiver.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (receiver *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	receiver.Index = uint(reader.ReadUint16())
}
