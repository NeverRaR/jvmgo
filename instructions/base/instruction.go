package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperation(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (receiver *NoOperandsInstruction) FetchOperation(reader *BytecodeReader) {
	//nothing to do
}

type BranchInstruction struct {
	offset int
}

func (receiver *BranchInstruction) FetchOperand(reader *BytecodeReader) {
	receiver.offset = int(reader.ReadInt16())
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
