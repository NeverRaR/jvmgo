package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
func (receiver *Stack) push(frame *Frame) {
	if receiver.size >= receiver.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if receiver._top != nil {
		frame.lower = receiver._top
	}
	receiver._top = frame
	receiver.size++
}

func (receiver *Stack) pop() *Frame {
	if receiver._top == nil {
		panic("jvm stack is empty!")
	}
	top := receiver._top
	receiver._top = top.lower
	top.lower = nil
	receiver.size--
	return top
}

func (receiver *Stack) top() *Frame {
	if receiver._top == nil {
		panic("jvm stack is empty!")
	}
	return receiver._top
}

func (receiver *Stack) isEmpty() bool {
	return receiver._top == nil
}

func (receiver *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, receiver.size)
	for frame := receiver._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}
