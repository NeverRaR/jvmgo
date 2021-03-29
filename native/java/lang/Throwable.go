package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {

}
