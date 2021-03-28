package misc

import (
	"jvmgo/instructions/base"
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize0)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}

// private static native void initialize();

func initialize0(frame *rtda.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)
	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)

	systemClass := vmClass.Loader().LoadClass("java/lang/System")
	setOut0Method := systemClass.GetStaticMethod("setOut0", "(Ljava/io/PrintStream;)V")
	printStreamClass := systemClass.Loader().LoadClass("java/io/PrintStream")
	frame.OperandStack().PushRef(printStreamClass.NewObject())
	base.InvokeMethod(frame, setOut0Method)
	//systemClass := vmClass.Loader().LoadClass("java/lang/System")
	//setOut0Method := systemClass.GetStaticMethod("setOut0", "(Ljava/io/PrintStream;)V")
	//newPrintStreamMethod := systemClass.GetStaticMethod("newPrintStream",
	//"(Ljava/io/FileOutputStream;Ljava/lang/String;)Ljava/io/PrintStream;")
	//fileOutputStreamClass := systemClass.Loader().LoadClass("java/io/FileOutputStream")
	//thread := frame.Thread()
	//newFrame := thread.NewFrame(setOut0Method)
	//thread.PushFrame(newFrame)
	//newFrame.OperandStack().PushRef(fileOutputStreamClass.NewObject())
	//newFrame.OperandStack().PushRef(nil)
	//base.InvokeMethod(newFrame, newPrintStreamMethod)

}
