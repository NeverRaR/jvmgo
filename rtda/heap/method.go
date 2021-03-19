package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocal     uint
	code         []byte
	argSlotCount uint
}

func (receiver *Method) ArgSlotCount() uint {
	return receiver.argSlotCount
}

func (receiver *Method) Code() []byte {
	return receiver.code
}

func (receiver *Method) MaxStack() uint {
	return receiver.maxStack
}

func (receiver *Method) MaxLocal() uint {
	return receiver.maxLocal
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (receiver *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(receiver.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		receiver.argSlotCount++
		if paramType == "J" || paramType == "D" {
			receiver.argSlotCount++
		}
	}
	if !receiver.IsStatic() {
		receiver.argSlotCount++ // `this` reference
	}
}

func (receiver *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		receiver.maxStack = codeAttr.MaxStack()
		receiver.maxLocal = codeAttr.MaxLocals()
		receiver.code = codeAttr.Code()
	}
}

func (receiver *Method) IsSynchronized() bool {
	return 0 != receiver.accessFlags&ACC_SYNCHRONIZED
}
func (receiver *Method) IsBridge() bool {
	return 0 != receiver.accessFlags&ACC_BRIDGE
}
func (receiver *Method) IsVarargs() bool {
	return 0 != receiver.accessFlags&ACC_VARARGS
}
func (receiver *Method) IsNative() bool {
	return 0 != receiver.accessFlags&ACC_NATIVE
}
func (receiver *Method) IsAbstract() bool {
	return 0 != receiver.accessFlags&ACC_ABSTRACT
}
func (receiver *Method) IsStrict() bool {
	return 0 != receiver.accessFlags&ACC_STRICT
}
