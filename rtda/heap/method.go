package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack                uint
	maxLocal                uint
	code                    []byte
	argSlotCount            uint
	exceptionTable          ExceptionTable
	exceptions              *classfile.ExceptionsAttribute // todo: rename
	parameterAnnotationData []byte                         // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte                         // AnnotationDefault_attribute
	parsedDescriptor        *MethodDescriptor
	lineNumberTable         *classfile.LineNumberTableAttribute
}

func (receiver *Method) LineNumberTable() *classfile.LineNumberTableAttribute {
	return receiver.lineNumberTable
}

func (receiver *Method) ParsedDescriptor() *MethodDescriptor {
	return receiver.parsedDescriptor
}

func (receiver *Method) AnnotationDefaultData() []byte {
	return receiver.annotationDefaultData
}

func (receiver *Method) ParameterAnnotationData() []byte {
	return receiver.parameterAnnotationData
}

func (receiver *Method) Exceptions() *classfile.ExceptionsAttribute {
	return receiver.exceptions
}

func (receiver *Method) GetLineNumber(pc int) int {
	if receiver.IsNative() {
		return -2
	}
	if receiver.lineNumberTable == nil {
		return -1
	}
	return receiver.lineNumberTable.GetLineNumber(pc)
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
	method.parsedDescriptor = md
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (receiver *Method) injectCodeAttribute(returnType string) {
	receiver.maxStack = 10
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
		receiver.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			receiver.class.constantPool)
		receiver.lineNumberTable = codeAttr.LineNumberTableAttribute()
	}
	receiver.exceptions = cfMethod.ExceptionsAttribute()
	receiver.annotationData = cfMethod.RuntimeVisibleAnnotationsAttributeData()
	receiver.parameterAnnotationData = cfMethod.RuntimeVisibleParameterAnnotationsAttributeData()
	receiver.annotationDefaultData = cfMethod.AnnotationDefaultAttributeData()
}

func (receiver *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := receiver.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
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
func (receiver *Method) isConstructor() bool {
	return !receiver.IsStatic() && receiver.name == "<init>"
}

// reflection
func (receiver *Method) ParameterTypes() []*Class {
	if receiver.argSlotCount == 0 {
		return nil
	}

	paramTypes := receiver.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = receiver.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}
func (receiver *Method) ReturnType() *Class {
	returnType := receiver.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return receiver.class.loader.LoadClass(returnClassName)
}
func (receiver *Method) ExceptionTypes() []*Class {
	if receiver.exceptions == nil {
		return nil
	}

	exIndexTable := receiver.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := receiver.class.constantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}
