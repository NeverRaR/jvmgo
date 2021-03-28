package heap

import "jvmgo/classfile"

type ClassMember struct {
	accessFlags    uint16
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
	class          *Class
}

func (receiver *ClassMember) AnnotationData() []byte {
	return receiver.annotationData
}

func (receiver *ClassMember) Signature() string {
	return receiver.signature
}

func (receiver *ClassMember) AccessFlags() uint16 {
	return receiver.accessFlags
}

func (receiver *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	receiver.accessFlags = memberInfo.AccessFlags()
	receiver.name = memberInfo.Name()
	receiver.descriptor = memberInfo.Descriptor()
}

func (receiver *ClassMember) IsPublic() bool {
	return 0 != receiver.accessFlags&ACC_PUBLIC
}
func (receiver *ClassMember) IsPrivate() bool {
	return 0 != receiver.accessFlags&ACC_PRIVATE
}
func (receiver *ClassMember) IsProtected() bool {
	return 0 != receiver.accessFlags&ACC_PROTECTED
}
func (receiver *ClassMember) IsStatic() bool {
	return 0 != receiver.accessFlags&ACC_STATIC
}
func (receiver *ClassMember) IsFinal() bool {
	return 0 != receiver.accessFlags&ACC_FINAL
}
func (receiver *ClassMember) IsSynthetic() bool {
	return 0 != receiver.accessFlags&ACC_SYNTHETIC
}

// jvms 5.4.4
func (receiver *ClassMember) isAccessibleTo(d *Class) bool {
	if receiver.IsPublic() {
		return true
	}
	c := receiver.class
	if receiver.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !receiver.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}

// getters
func (receiver *ClassMember) Name() string {
	return receiver.name
}
func (receiver *ClassMember) Descriptor() string {
	return receiver.descriptor
}
func (receiver *ClassMember) Class() *Class {
	return receiver.class
}
