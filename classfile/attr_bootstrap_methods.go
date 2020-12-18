package classfile

type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (receiver *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	receiver.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range receiver.bootstrapMethods {
		receiver.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
