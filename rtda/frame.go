package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func NewFrame(maxLocal, maxStack uint) *Frame {
	return &Frame{
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
