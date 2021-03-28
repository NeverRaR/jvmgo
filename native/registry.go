package native

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {

}

func initSystem(frame *rtda.Frame) {
	systemClass := frame.Method().Class()
	setOut0Method := systemClass.GetStaticMethod("setOut0", "(Ljava/io/PrintStream;)V")
	fileOutputStreamClass := systemClass.Loader().LoadClass("java/io/FileOutputStream")
	frame.OperandStack().PushRef(fileOutputStreamClass.NewObject())
	base.InvokeMethod(frame, setOut0Method)
}
