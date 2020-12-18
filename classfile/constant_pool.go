package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (receiver ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := receiver[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (receiver ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := receiver.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := receiver.getUtf8(ntInfo.nameIndex)
	_type := receiver.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (receiver ConstantPool) getClassName(index uint16) string {
	classInfo := receiver.getConstantInfo(index).(*ConstantClassInfo)
	return receiver.getUtf8(classInfo.nameIndex)
}

func (receiver ConstantPool) getUtf8(index uint16) string {
	utf8Info := receiver.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
