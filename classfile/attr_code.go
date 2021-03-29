package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (receiver *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return receiver.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (e ExceptionTableEntry) CatchType() uint16 {
	return e.catchType
}

func (e ExceptionTableEntry) HandlerPc() uint16 {
	return e.handlerPc
}

func (e ExceptionTableEntry) EndPc() uint16 {
	return e.endPc
}

func (e ExceptionTableEntry) StartPc() uint16 {
	return e.startPc
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
