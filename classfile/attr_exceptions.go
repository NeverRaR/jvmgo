package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (receiver *ExceptionsAttribute) readInfo(reader *ClassReader) {
	receiver.exceptionIndexTable = reader.readUint16s()
}
func (receiver *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return receiver.exceptionIndexTable
}
