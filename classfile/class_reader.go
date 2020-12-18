package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (receiver *ClassReader) readUint8() uint8 {
	val := receiver.data[0]
	receiver.data = receiver.data[1:]
	return val
}

func (receiver *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(receiver.data)
	receiver.data = receiver.data[2:]
	return val
}

func (receiver *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(receiver.data)
	receiver.data = receiver.data[4:]
	return val
}

func (receiver *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(receiver.data)
	receiver.data = receiver.data[8:]
	return val
}

func (receiver *ClassReader) readUint16s() []uint16 {
	n := receiver.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = receiver.readUint16()
	}
	return s
}

func (receiver *ClassReader) readBytes(n uint32) []byte {
	bytes := receiver.data[:n]
	receiver.data = receiver.data[n:]
	return bytes
}
