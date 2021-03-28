package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (receiver *JVM) start() {
	receiver.initVM()
	receiver.execMain()
}

func (receiver *JVM) initVM() {
	vmClass := receiver.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(receiver.mainThread, vmClass)
	interpret(receiver.mainThread, receiver.cmd.verboseInstFlag)
}

func (receiver *JVM) execMain() {
	className := strings.Replace(receiver.cmd.class, ".", "/", -1)
	mainClass := receiver.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", receiver.cmd.class)
		return
	}
	argsArr := receiver.createArgsArray()
	frame := receiver.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	receiver.mainThread.PushFrame(frame)
	interpret(receiver.mainThread, receiver.cmd.verboseInstFlag)
}

func (receiver *JVM) createArgsArray() *heap.Object {
	stringClass := receiver.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(receiver.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range receiver.cmd.args {
		jArgs[i] = heap.JString(receiver.classLoader, arg)
	}
	return argsArr
}
