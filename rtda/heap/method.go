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
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (receiver *Method) injectCodeAttribute(returnType string) {
	receiver.maxStack = 4
	receiver.maxLocal = receiver.argSlotCount
	switch returnType[0] {
	case 'V':
		receiver.code = []byte{0xfe, 0xb1} //return
	case 'D':
		receiver.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		receiver.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		receiver.code = []byte{0xfe, 0xad} //lreturn
	case 'L', '[':
		receiver.code = []byte{0xfe, 0xb0} //areturn
	default:
		receiver.code = []byte{0xfe, 0xac} //ireturn

	}
}

func (receiver *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
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
