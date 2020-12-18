package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (receiver *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	receiver.nameIndex = reader.readUint16()
	receiver.descriptorIndex = reader.readUint16()
}
