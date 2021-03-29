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
	//init savedProps
	thread := frame.Thread()
	vmClass := frame.Method().Class()
	systemClass := vmClass.Loader().LoadClass("java/lang/System")
	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	ops := rtda.NewOperandStack(8)
	shimFrame := rtda.NewShimFrame(thread, ops)

	//init savedProps
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "java.lang.Integer.IntegerCache.high")
	val := heap.JString(vmClass.Loader(), "127")
	ops.PushRef(savedProps)
	ops.PushRef(key)
	ops.PushRef(val)
	thread.PushFrame(shimFrame)
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(shimFrame, setPropMethod)

	//init props
	ops = rtda.NewOperandStack(8)
	initPropertiesMethod := systemClass.GetStaticMethod("initProperties",
		"(Ljava/util/Properties;)Ljava/util/Properties;")
	props := propsClass.NewObject()
	ops.PushRef(props)
	shimFrame = rtda.NewShimFrame(thread, ops)
	thread.PushFrame(shimFrame)
	base.InvokeMethod(shimFrame, initPropertiesMethod)
	ops = rtda.NewOperandStack(8)
	propsConstructor := propsClass.GetConstructor("()V")
	ops.PushRef(props)
	shimFrame = rtda.NewShimFrame(thread, ops)
	base.InvokeMethod(shimFrame, propsConstructor)
	systemClass.SetRefVar("props", "Ljava/util/Properties;", props)

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
