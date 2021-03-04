package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (receiver *CodeAttribute) readInfo(reader *ClassReader) {
	receiver.maxStack = reader.readUint16()
	receiver.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	receiver.code = reader.readBytes(codeLength)
	receiver.exceptionTable = readExceptionTable(reader)
	receiver.attributes = readAttributes(reader, receiver.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (receiver *CodeAttribute) MaxLocals() uint {
	return uint(receiver.maxLocals)
}

func (receiver *CodeAttribute) MaxStack() uint {
	return uint(receiver.maxStack)
}

func (receiver *CodeAttribute) Code() []byte {
	return receiver.code
}
