package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func (receiver *Frame) SetNextPC(nextPC int) {
	receiver.nextPC = nextPC
}

func (receiver *Frame) NextPC() int {
	return receiver.nextPC
}

func (receiver *Frame) Thread() *Thread {
	return receiver.thread
}

func newFrame(thread *Thread, maxLocal, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocal),
		operandStack: newOperandStack(maxStack),
	}
}

func (receiver *Frame) LocalVars() LocalVars {
	return receiver.localVars
}

func (receiver *Frame) OperandStack() *OperandStack {
	return receiver.operandStack
}
