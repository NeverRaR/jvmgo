package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func (receiver *Thread) PC() int {
	return receiver.pc
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
func (receiver *Thread) SetPC(pc int) {
	receiver.pc = pc
}
func (receiver *Thread) PushFrame(frame *Frame) {
	receiver.stack.push(frame)
}
func (receiver *Thread) PopFrame() *Frame {
	return receiver.stack.pop()
}
func (receiver *Thread) CurrentFrame() *Frame {
	return receiver.stack.top()
}

func (receiver *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(receiver, maxLocals, maxStack)
}
