package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"runtime"
)

func init() {
	native.Register("java/lang/Runtime", "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OperandStack()
	stack.PushInt(int32(numCPU))
}
