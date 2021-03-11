package rtda

import "jvmgo/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
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

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocal()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (receiver *Frame) LocalVars() LocalVars {
	return receiver.localVars
}

func (receiver *Frame) Method() *heap.Method {
	return receiver.method
}

func (receiver *Frame) OperandStack() *OperandStack {
	return receiver.operandStack
}
