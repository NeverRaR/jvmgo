package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (receiver *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	receiver.referenceKind = reader.readUint8()
	receiver.referenceIndex = reader.readUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (receiver *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	receiver.descriptorIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (receiver *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	receiver.bootstrapMethodAttrIndex = reader.readUint16()
	receiver.nameAndTypeIndex = reader.readUint16()
}
