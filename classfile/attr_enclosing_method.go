package classfile

type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (receiver *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	receiver.classIndex = reader.readUint16()
	receiver.methodIndex = reader.readUint16()
}

func (receiver *EnclosingMethodAttribute) ClassName() string {
	return receiver.cp.getClassName(receiver.classIndex)
}

func (receiver *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if receiver.methodIndex > 0 {
		return receiver.cp.getNameAndType(receiver.methodIndex)
	} else {
		return "", ""
	}
}
