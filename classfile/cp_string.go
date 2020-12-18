package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (receiver *ConstantStringInfo) readInfo(reader *ClassReader) {
	receiver.stringIndex = reader.readUint16()
}

func (receiver *ConstantStringInfo) String() string {
	return receiver.cp.getUtf8(receiver.stringIndex)
}
