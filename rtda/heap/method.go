package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocal uint
	code     []byte
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
	}
	return methods
}

func (receiver *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		receiver.maxStack = codeAttr.MaxStack()
		receiver.maxLocal = codeAttr.MaxLocals()
		receiver.code = codeAttr.Code()
	}
}
