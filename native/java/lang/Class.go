package lang

import (
	"jvmgo/instructions/base"
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"strings"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register("java/lang/Class", "isInterface",
		"()Z", isInterface)
	native.Register("java/lang/Class", "isPrimitive",
		"()Z", isPrimitive)
	native.Register("java/lang/Class", "getDeclaredFields0",
		"(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native.Register("java/lang/Class", "forName0",
		"(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.Register("java/lang/Class", "getDeclaredConstructors0",
		"(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.Register("java/lang/Class", "getModifiers",
		"()I", getModifiers)
	native.Register("java/lang/Class", "getSuperclass",
		"()Ljava/lang/Class;", getSuperclass)
	native.Register("java/lang/Class", "getInterfaces0",
		"()[Ljava/lang/Class;", getInterfaces0)
	native.Register("java/lang/Class", "isArray",
		"()Z", isArray)
	native.Register("java/lang/Class", "getDeclaredMethods0",
		"(Z)[Ljava/lang/reflect/Method;", getDeclaredMethods0)
	native.Register("java/lang/Class", "getComponentType",
		"()Ljava/lang/Class;", getComponentType)
	native.Register("java/lang/Class", "isAssignableFrom",
		"(Ljava/lang/Class;)Z", isAssignableFrom)
	native.Register("java/lang/Class", "getEnclosingMethod0",
		"()[Ljava/lang/Object;", getEnclosingMethod0)
	native.Register("java/lang/Class", "getDeclaringClass0",
		"()Ljava/lang/Class;", getDeclaringClass0)
}

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}

// private native String getName0();
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}

// private static native Class<?> forName0(String name, boolean initialize,
//                                         ClassLoader loader,
//                                         Class<?> caller)
//     throws ClassNotFoundException;
// (Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	//jLoader := vars.GetRef(2)

	goName := heap.GoString(jName)
	goName = strings.Replace(goName, ".", "/", -1)
	goClass := frame.Method().Class().Loader().LoadClass(goName)
	jClass := goClass.JClass()

	if initialize && !goClass.InitStarted() {
		// undo forName0
		thread := frame.Thread()
		frame.SetNextPC(thread.PC())
		// init class
		base.InitClass(thread, goClass)
	} else {
		stack := frame.OperandStack()
		stack.PushRef(jClass)
	}
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	modifiers := class.AccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(modifiers))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	superClass := class.SuperClass()

	stack := frame.OperandStack()
	if superClass != nil {
		stack.PushRef(superClass.JClass())
	} else {
		stack.PushRef(nil)
	}
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	interfaces := class.Interfaces()
	classArr := toClassArr(class.Loader(), interfaces)

	stack := frame.OperandStack()
	stack.PushRef(classArr)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsArray())
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass()

	stack := frame.OperandStack()
	stack.PushRef(componentClassObj)
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	cls := vars.GetRef(1)

	thisClass := this.Extra().(*heap.Class)
	clsClass := cls.Extra().(*heap.Class)
	ok := thisClass.IsAssignableFrom(clsClass)

	stack := frame.OperandStack()
	stack.PushBoolean(ok)
}

const _fieldConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"

// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	classLoader := frame.Method().Class().Loader()
	fieldClass := classLoader.LoadClass("java/lang/reflect/Field")
	fieldArr := fieldClass.ArrayClass().NewArray(fieldCount)

	stack := frame.OperandStack()
	stack.PushRef(fieldArr)

	if fieldCount > 0 {
		thread := frame.Thread()
		fieldObjs := fieldArr.Refs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObject()
			fieldObj.SetExtra(goField)
			fieldObjs[i] = fieldObj

			ops := rtda.NewOperandStack(8)
			ops.PushRef(fieldObj)                                          // this
			ops.PushRef(classObj)                                          // declaringClass
			ops.PushRef(heap.JString(classLoader, goField.Name()))         // name
			ops.PushRef(goField.Type().JClass())                           // type
			ops.PushInt(int32(goField.AccessFlags()))                      // modifiers
			ops.PushInt(int32(goField.SlotId()))                           // slot
			ops.PushRef(getSignatureStr(classLoader, goField.Signature())) // signature
			ops.PushRef(toByteArr(classLoader, goField.AnnotationData()))  // annotations

			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)

			// init fieldObj
			base.InvokeMethod(shimFrame, fieldConstructor)
		}
	}
}

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	classLoader := frame.Method().Class().Loader()
	constructorClass := classLoader.LoadClass("java/lang/reflect/Constructor")
	constructorArr := constructorClass.ArrayClass().NewArray(constructorCount)

	stack := frame.OperandStack()
	stack.PushRef(constructorArr)

	if constructorCount > 0 {
		thread := frame.Thread()
		constructorObjs := constructorArr.Refs()
		constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
		for i, constructor := range constructors {
			constructorObj := constructorClass.NewObject()
			constructorObj.SetExtra(constructor)
			constructorObjs[i] = constructorObj

			ops := rtda.NewOperandStack(9)
			ops.PushRef(constructorObj)                                                // this
			ops.PushRef(classObj)                                                      // declaringClass
			ops.PushRef(toClassArr(classLoader, constructor.ParameterTypes()))         // parameterTypes
			ops.PushRef(toClassArr(classLoader, constructor.ExceptionTypes()))         // checkedExceptions
			ops.PushInt(int32(constructor.AccessFlags()))                              // modifiers
			ops.PushInt(int32(0))                                                      // todo slot
			ops.PushRef(getSignatureStr(classLoader, constructor.Signature()))         // signature
			ops.PushRef(toByteArr(classLoader, constructor.AnnotationData()))          // annotations
			ops.PushRef(toByteArr(classLoader, constructor.ParameterAnnotationData())) // parameterAnnotations

			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)

			// init constructorObj
			base.InvokeMethod(shimFrame, constructorInitMethod)
		}
	}
}

/*
Method(Class<?> declaringClass,
       String name,
       Class<?>[] parameterTypes,
       Class<?> returnType,
       Class<?>[] checkedExceptions,
       int modifiers,
       int slot,
       String signature,
       byte[] annotations,
       byte[] parameterAnnotations,
       byte[] annotationDefault)
*/
const _methodConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"[Ljava/lang/Class;" +
	"Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B[B)V"

// private native Method[] getDeclaredMethods0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Method;
func getDeclaredMethods0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))

	classLoader := class.Loader()
	methodClass := classLoader.LoadClass("java/lang/reflect/Method")
	methodArr := methodClass.ArrayClass().NewArray(methodCount)

	stack := frame.OperandStack()
	stack.PushRef(methodArr)

	// create method objs
	if methodCount > 0 {
		thread := frame.Thread()
		methodObjs := methodArr.Refs()
		methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
		for i, method := range methods {
			methodObj := methodClass.NewObject()
			methodObj.SetExtra(method)
			methodObjs[i] = methodObj

			ops := rtda.NewOperandStack(12)
			ops.PushRef(methodObj)                                                // this
			ops.PushRef(classObj)                                                 // declaringClass
			ops.PushRef(heap.JString(classLoader, method.Name()))                 // name
			ops.PushRef(toClassArr(classLoader, method.ParameterTypes()))         // parameterTypes
			ops.PushRef(method.ReturnType().JClass())                             // returnType
			ops.PushRef(toClassArr(classLoader, method.ExceptionTypes()))         // checkedExceptions
			ops.PushInt(int32(method.AccessFlags()))                              // modifiers
			ops.PushInt(int32(0))                                                 // todo: slot
			ops.PushRef(getSignatureStr(classLoader, method.Signature()))         // signature
			ops.PushRef(toByteArr(classLoader, method.AnnotationData()))          // annotations
			ops.PushRef(toByteArr(classLoader, method.ParameterAnnotationData())) // parameterAnnotations
			ops.PushRef(toByteArr(classLoader, method.AnnotationDefaultData()))   // annotationDefault

			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)

			// init methodObj
			base.InvokeMethod(shimFrame, methodConstructor)
		}
	}
}
func getEnclosingMethod0(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func getDeclaringClass0(frame *rtda.Frame) {
	//vars := frame.LocalVars()
	//classObj := vars.GetThis()
	frame.OperandStack().PushRef(nil)
}
