package classfile

import "fmt"

type ClassFile struct {
	//magic	uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (receiver *ClassFile) ConstantPool() ConstantPool {
	return receiver.constantPool
}

func (receiver *ClassFile) Methods() []*MemberInfo {
	return receiver.methods
}

func (receiver *ClassFile) Fields() []*MemberInfo {
	return receiver.fields
}

func (receiver *ClassFile) AccessFlags() uint16 {
	return receiver.accessFlags
}

func (receiver *ClassFile) MinorVersion() uint16 {
	return receiver.minorVersion
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (receiver *ClassFile) read(reader *ClassReader) {
	receiver.readAndCheckMagic(reader)
	receiver.readAndCheckVersion(reader)
	receiver.constantPool = readConstantPool(reader)
	receiver.accessFlags = reader.readUint16()
	receiver.thisClass = reader.readUint16()
	receiver.superClass = reader.readUint16()
	receiver.interfaces = reader.readUint16s()
	receiver.fields = readMembers(reader, receiver.constantPool)
	receiver.methods = readMembers(reader, receiver.constantPool)
	receiver.attributes = readAttributes(reader, receiver.constantPool)
}

func (receiver *ClassFile) MajorVersion() uint16 {
	return receiver.majorVersion
}

func (receiver *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (receiver *ClassFile) readAndCheckVersion(reader *ClassReader) {
	receiver.minorVersion = reader.readUint16()
	receiver.majorVersion = reader.readUint16()
	switch receiver.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if receiver.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (receiver *ClassFile) ClassName() string {
	return receiver.constantPool.getClassName(receiver.thisClass)
}

func (receiver *ClassFile) SuperClassName() string {
	if receiver.superClass > 0 {
		return receiver.constantPool.getClassName(receiver.superClass)
	}
	return ""
}

func (receiver *ClassFile) InterfaceNames() []string {
	interfacesName := make([]string, len(receiver.interfaces))
	for i, cpIndex := range receiver.interfaces {
		interfacesName[i] = receiver.constantPool.getClassName(cpIndex)
	}
	return interfacesName
}
