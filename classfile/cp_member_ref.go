package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (receiver *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	receiver.classIndex = reader.readUint16()
	receiver.nameAndTypeIndex = reader.readUint16()
}

func (receiver *ConstantMemberRefInfo) ClassName() string {
	return receiver.cp.getClassName(receiver.classIndex)
}

func (receiver *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return receiver.cp.getNameAndType(receiver.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }
type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }
type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }
