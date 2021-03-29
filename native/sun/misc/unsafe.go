package misc

import (
	"encoding/binary"
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"math"
)

func init() {
	native.Register("sun/misc/Unsafe", "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	native.Register("sun/misc/Unsafe", "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	native.Register("sun/misc/Unsafe", "addressSize", "()I", addressSize)
	native.Register("sun/misc/Unsafe", "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	native.Register("sun/misc/Unsafe", "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	native.Register("sun/misc/Unsafe", "getIntVolatile", "(Ljava/lang/Object;J)I", getInt)
	native.Register("sun/misc/Unsafe", "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	native.Register("sun/misc/Unsafe", "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;", getObject)
	native.Register("sun/misc/Unsafe", "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z", compareAndSwapLong)
	native.Register("sun/misc/Unsafe", "allocateMemory", "(J)J", allocateMemory)
	native.Register("sun/misc/Unsafe", "reallocateMemory", "(JJ)J", reallocateMemory)
	native.Register("sun/misc/Unsafe", "freeMemory", "(J)V", freeMemory)
	native.Register("sun/misc/Unsafe", "getByte", "(J)B", mem_getByte)
	native.Register("sun/misc/Unsafe", "putLong", "(JJ)V", mem_putLong)

}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(1) // todo
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
	// vars := frame.LocalVars()
	// vars.GetRef(0) // this

	stack := frame.OperandStack()
	stack.PushInt(8) // todo unsafe.Sizeof(int)
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jField := vars.GetRef(1)

	offset := jField.GetIntVar("slot", "I")

	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public final native boolean compareAndSwapObject(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSwapObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	fields := obj.Data()
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVal := vars.GetRef(5)

	// todo
	if anys, ok := fields.(heap.Slots); ok {
		// object
		swapped := _casObj(obj, anys, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		swapped := _casArr(objs, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapObject!")
	}
}
func _casObj(obj *heap.Object, fields heap.Slots, offset int64, expected, newVal *heap.Object) bool {
	current := fields.GetRef(uint(offset))
	if current == expected {
		fields.SetRef(uint(offset), newVal)
		return true
	} else {
		return false
	}
}
func _casArr(objs []*heap.Object, offset int64, expected, newVal *heap.Object) bool {
	current := objs[offset]
	if current == expected {
		objs[offset] = newVal
		return true
	} else {
		return false
	}
}

// public native boolean getInt(Object o, long offset);
// (Ljava/lang/Object;J)I
func getInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if slots, ok := fields.(heap.Slots); ok {
		// object
		stack.PushInt(slots.GetInt(uint(offset)))
	} else if shorts, ok := fields.([]int32); ok {
		// int[]
		stack.PushInt(int32(shorts[offset]))
	} else {
		panic("getInt!")
	}
}

// public final native boolean compareAndSwapInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	newVal := vars.GetInt(5)

	if slots, ok := fields.(heap.Slots); ok {
		// object
		oldVal := slots.GetInt(uint(offset))
		if oldVal == expected {
			slots.SetInt(uint(offset), newVal)
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else if ints, ok := fields.([]int32); ok {
		// int[]
		oldVal := ints[offset]
		if oldVal == expected {
			ints[offset] = newVal
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else {
		// todo
		panic("todo: compareAndSwapInt!")
	}
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)

	if anys, ok := fields.(heap.Slots); ok {
		// object
		x := anys.GetRef(uint(offset))
		frame.OperandStack().PushRef(x)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		x := objs[offset]
		frame.OperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}

// public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSwapLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)
	expected := vars.GetLong(4)
	newVal := vars.GetLong(6)

	if slots, ok := fields.(heap.Slots); ok {
		// object
		oldVal := slots.GetLong(uint(offset))
		if oldVal == expected {
			slots.SetLong(uint(offset), newVal)
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else if longs, ok := fields.([]int64); ok {
		// long[]
		oldVal := longs[offset]
		if oldVal == expected {
			longs[offset] = newVal
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else {
		// todo
		panic("todo: compareAndSwapLong!")
	}
}

// public native long allocateMemory(long bytes);
// (J)J
func allocateMemory(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	bytes := vars.GetLong(1)

	address := allocate(bytes)
	stack := frame.OperandStack()
	stack.PushLong(address)
}

// public native long reallocateMemory(long address, long bytes);
// (JJ)J
func reallocateMemory(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	address := vars.GetLong(1)
	bytes := vars.GetLong(3)

	newAddress := reallocate(address, bytes)
	stack := frame.OperandStack()
	stack.PushLong(newAddress)
}

// public native void freeMemory(long address);
// (J)V
func freeMemory(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	address := vars.GetLong(1)
	free(address)
}

//// public native void putAddress(long address, long x);
//// (JJ)V
//func putAddress(frame *rtda.Frame) {
//	mem_putLong(frame)
//}
//
//// public native long getAddress(long address);
//// (J)J
//func getAddress(frame *rtda.Frame) {
//	mem_getLong(frame)
//}
//
//// public native void putByte(long address, byte x);
//// (JB)V
//func mem_putByte(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutInt8(mem, int8(value.(int32)))
//}
//
// public native byte getByte(long address);
// (J)B
func mem_getByte(frame *rtda.Frame) {
	stack, mem := _get(frame)
	stack.PushInt(int32(Int8(mem)))
}

//
//// public native void putShort(long address, short x);
//// (JS)V
//func mem_putShort(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutInt16(mem, int16(value.(int32)))
//}
//
//// public native short getShort(long address);
//// (J)S
//func mem_getShort(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushInt(int32(Int16(mem)))
//}
//
//// public native void putChar(long address, char x);
//// (JC)V
//func mem_putChar(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutUint16(mem, uint16(value.(int32)))
//}
//
//// public native char getChar(long address);
//// (J)C
//func mem_getChar(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushInt(int32(Uint16(mem)))
//}
//
//// public native void putInt(long address, int x);
//// (JI)V
//func mem_putInt(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutInt32(mem, value.(int32))
//}
//
//// public native int getInt(long address);
//// (J)I
//func mem_getInt(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushInt(Int32(mem))
//}
//
// public native void putLong(long address, long x);
// (JJ)V
func mem_putLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	address := vars.GetLong(1)
	value := vars.GetLong(3)

	mem := memoryAt(address)

	PutInt64(mem, value)
}

//
//// public native long getLong(long address);
//// (J)J
//func mem_getLong(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushLong(Int64(mem))
//}
//
//// public native void putFloat(long address, float x);
//// (JJ)V
//func mem_putFloat(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutFloat32(mem, value.(float32))
//}
//
//// public native float getFloat(long address);
//// (J)J
//func mem_getFloat(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushFloat(Float32(mem))
//}
//
//// public native void putDouble(long address, double x);
//// (JJ)V
//func mem_putDouble(frame *rtda.Frame) {
//	mem, value := _put(frame)
//	PutFloat64(mem, value.(float64))
//}
//
//// public native double getDouble(long address);
//// (J)J
//func mem_getDouble(frame *rtda.Frame) {
//	stack, mem := _get(frame)
//	stack.PushDouble(Float64(mem))
//}
//
//func _put(frame *rtda.Frame) ([]byte, interface{}) {
//	vars := frame.LocalVars()
//	// vars.GetRef(0) // this
//	address := vars.GetLong(1)
//	value := vars.Get(3)
//
//	mem := memoryAt(address)
//	return mem, value
//}
//
func _get(frame *rtda.Frame) (*rtda.OperandStack, []byte) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	address := vars.GetLong(1)

	stack := frame.OperandStack()
	mem := memoryAt(address)
	return stack, mem
}

var _bigEndian = binary.BigEndian

func PutInt8(s []byte, val int8) {
	s[0] = uint8(val)
}
func Int8(s []byte) int8 {
	return int8(s[0])
}

func PutUint16(s []byte, val uint16) {
	_bigEndian.PutUint16(s, val)
}
func Uint16(s []byte) uint16 {
	return _bigEndian.Uint16(s)
}

func PutInt16(s []byte, val int16) {
	_bigEndian.PutUint16(s, uint16(val))
}
func Int16(s []byte) int16 {
	return int16(_bigEndian.Uint16(s))
}

func PutInt32(s []byte, val int32) {
	_bigEndian.PutUint32(s, uint32(val))
}
func Int32(s []byte) int32 {
	return int32(_bigEndian.Uint32(s))
}

func PutInt64(s []byte, val int64) {
	_bigEndian.PutUint64(s, uint64(val))
}
func Int64(s []byte) int64 {
	return int64(_bigEndian.Uint64(s))
}

func PutFloat32(s []byte, val float32) {
	_bigEndian.PutUint32(s, math.Float32bits(val))
}
func Float32(s []byte) float32 {
	return math.Float32frombits(_bigEndian.Uint32(s))
}

func PutFloat64(s []byte, val float64) {
	_bigEndian.PutUint64(s, math.Float64bits(val))
}
func Float64(s []byte) float64 {
	return math.Float64frombits(_bigEndian.Uint64(s))
}
