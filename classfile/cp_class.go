package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (receiver *ConstantClassInfo) readInfo(reader *ClassReader) {
	receiver.nameIndex = reader.readUint16()
}

func (receiver *ConstantClassInfo) Name() string {
	return receiver.cp.getUtf8(receiver.nameIndex)
}
