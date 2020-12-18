package classfile

type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (receiver *SignatureAttribute) readInfo(reader *ClassReader) {
	receiver.signatureIndex = reader.readUint16()
}

func (receiver *SignatureAttribute) Signature() string {
	return receiver.cp.getUtf8(receiver.signatureIndex)
}
