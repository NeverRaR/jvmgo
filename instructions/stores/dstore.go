package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

// Store double into local variable
type DSTORE struct{ base.Index8Instruction }

func (receiver *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(receiver.Index))
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (receiver *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (receiver *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (receiver *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (receiver *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
