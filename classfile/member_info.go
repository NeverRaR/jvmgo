package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (receiver *MemberInfo) AccessFlags() uint16 {
	return receiver.accessFlags
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (receiver *MemberInfo) Name() string {
	return receiver.cp.getUtf8(receiver.nameIndex)
}

func (receiver *MemberInfo) Descriptor() string {
	return receiver.cp.getUtf8(receiver.descriptorIndex)
}

func (receiver *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range receiver.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (receiver *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range receiver.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (receiver *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range receiver.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (receiver *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return receiver.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (receiver *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return receiver.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (receiver *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return receiver.getUnparsedAttributeData("AnnotationDefault")
}

func (receiver *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range receiver.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
