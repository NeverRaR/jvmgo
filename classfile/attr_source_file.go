package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (receiver *SourceFileAttribute) readInfo(reader *ClassReader) {
	receiver.sourceFileIndex = reader.readUint16()
}
func (receiver *SourceFileAttribute) FileName() string {
	return receiver.cp.getUtf8(receiver.sourceFileIndex)
}
