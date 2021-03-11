package heap

import "jvmgo/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (receiver *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		receiver.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (receiver *Field) IsVolatile() bool {
	return 0 != receiver.accessFlags&ACC_VOLATILE
}
func (receiver *Field) IsTransient() bool {
	return 0 != receiver.accessFlags&ACC_TRANSIENT
}
func (receiver *Field) IsEnum() bool {
	return 0 != receiver.accessFlags&ACC_ENUM
}

func (receiver *Field) ConstValueIndex() uint {
	return receiver.constValueIndex
}
func (receiver *Field) SlotId() uint {
	return receiver.slotId
}
func (receiver *Field) isLongOrDouble() bool {
	return receiver.descriptor == "J" || receiver.descriptor == "D"
}
