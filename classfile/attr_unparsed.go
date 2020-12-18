package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (receiver *UnparsedAttribute) readInfo(reader *ClassReader) {
	receiver.info = reader.readBytes(receiver.length)
}
