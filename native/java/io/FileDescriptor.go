package io

import (
	"jvmgo/native"
	"jvmgo/rtda"
)

func init() {
	native.Register("java/io/FileDescriptor", "set", "(I)J", set)
}

// private static native long set(int d);
// (I)J
func set(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
