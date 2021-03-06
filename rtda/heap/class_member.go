package heap

import "jvmgo/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
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
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !receiver.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
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
