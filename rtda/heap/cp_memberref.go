package heap

import "jvmgo/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (receiver *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	receiver.className = refInfo.ClassName()
	receiver.name, receiver.descriptor = refInfo.NameAndDescriptor()
}

func (receiver *MemberRef) Name() string {
	return receiver.name
}
func (receiver *MemberRef) Descriptor() string {
	return receiver.descriptor
}
