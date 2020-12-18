package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

func (receiver *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	receiver.val = int32(bytes)
}

func (receiver *ConstantIntegerInfo) Val() int32 {
	return receiver.val
}

type ConstantFloatInfo struct {
	val float32
}

func (receiver *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	receiver.val = math.Float32frombits(bytes)
}

func (receiver *ConstantFloatInfo) Val() float32 {
	return receiver.val
}

type ConstantLongInfo struct {
	val int64
}

func (receiver *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	receiver.val = int64(bytes)
}

func (receiver *ConstantLongInfo) Val() int64 {
	return receiver.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (receiver *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	receiver.val = math.Float64frombits(bytes)
}

func (receiver *ConstantDoubleInfo) Val() float64 {
	return receiver.val
}
