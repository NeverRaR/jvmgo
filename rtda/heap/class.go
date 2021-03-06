package heap

import (
	"jvmgo/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (receiver *Class) IsPublic() bool {
	return 0 != receiver.accessFlags&ACC_PUBLIC
}
func (receiver *Class) IsFinal() bool {
	return 0 != receiver.accessFlags&ACC_FINAL
}
func (receiver *Class) IsSuper() bool {
	return 0 != receiver.accessFlags&ACC_SUPER
}
func (receiver *Class) IsInterface() bool {
	return 0 != receiver.accessFlags&ACC_INTERFACE
}
func (receiver *Class) IsAbstract() bool {
	return 0 != receiver.accessFlags&ACC_ABSTRACT
}
func (receiver *Class) IsSynthetic() bool {
	return 0 != receiver.accessFlags&ACC_SYNTHETIC
}
func (receiver *Class) IsAnnotation() bool {
	return 0 != receiver.accessFlags&ACC_ANNOTATION
}
func (receiver *Class) IsEnum() bool {
	return 0 != receiver.accessFlags&ACC_ENUM
}

// getters
func (receiver *Class) Name() string {
	return receiver.name
}
func (receiver *Class) ConstantPool() *ConstantPool {
	return receiver.constantPool
}
func (receiver *Class) Fields() []*Field {
	return receiver.fields
}
func (receiver *Class) Methods() []*Method {
	return receiver.methods
}
func (receiver *Class) SuperClass() *Class {
	return receiver.superClass
}
func (receiver *Class) StaticVars() Slots {
	return receiver.staticVars
}

func (receiver *Class) isAccessibleTo(other *Class) bool {
	return receiver.IsPublic() ||
		receiver.getPackageName() == other.getPackageName()
}

func (receiver *Class) getPackageName() string {
	if i := strings.LastIndex(receiver.name, "/"); i >= 0 {
		return receiver.name[:i]
	}
	return ""
}

func (receiver *Class) NewObject() *Object {
	return newObject(receiver)
}

func (receiver *Class) GetMainMethod() *Method {
	return receiver.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (receiver *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range receiver.methods {
		if method.IsStatic() && method.name == name && descriptor == method.descriptor {
			return method
		}
	}
	return nil
}
