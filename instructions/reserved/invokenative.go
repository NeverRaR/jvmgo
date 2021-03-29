package reserved

import (
	"jvmgo/instructions/base"
	"jvmgo/native"
	"jvmgo/rtda"
)

import _ "jvmgo/native/java/lang"
import _ "jvmgo/native/sun/misc"
import _ "jvmgo/native/sun/reflect"
import _ "jvmgo/native/java/security"
import _ "jvmgo/native/java/io"
import _ "jvmgo/native/sun/io"

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (receiver *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)

	}
	nativeMethod(frame)
}
