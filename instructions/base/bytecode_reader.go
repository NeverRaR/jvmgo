package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (receiver *BytecodeReader) Reset(code []byte, pc int) {
	receiver.code = code
	receiver.pc = pc
}
func (receiver *BytecodeReader) ReadUint8() uint8 {
	i := receiver.code[receiver.pc]
	receiver.pc++
	return i
}
func (receiver *BytecodeReader) ReadInt8() int8 {
	return int8(receiver.ReadUint8())
}
func (receiver *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(receiver.ReadUint8())
	byte2 := uint16(receiver.ReadUint8())
	return (byte1 << 8) | byte2
}
func (receiver *BytecodeReader) ReadInt16() int16 {
	return int16(receiver.ReadUint16())
}

func (receiver *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(receiver.ReadUint8())
	byte2 := int32(receiver.ReadUint8())
	byte3 := int32(receiver.ReadUint8())
	byte4 := int32(receiver.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (receiver *BytecodeReader) SkipPadding() {
	for receiver.pc%4 != 0 {
		receiver.ReadUint8()
	}
}

func (receiver *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = receiver.ReadInt32()
	}
	return ints
}
