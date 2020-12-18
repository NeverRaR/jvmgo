package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (receiver *ConstantValueAttribute) readInfo(reader *ClassReader) {
	receiver.constantValueIndex = reader.readUint16()
}
func (receiver *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return receiver.constantValueIndex
}
