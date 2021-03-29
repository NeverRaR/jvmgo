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
	sourceFile        string
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
}

func (receiver *Class) SourceFile() string {
	return receiver.sourceFile
}

func (receiver *Class) Interfaces() []*Class {
	return receiver.interfaces
}

func (receiver *Class) AccessFlags() uint16 {
	return receiver.accessFlags
}

func (receiver *Class) JClass() *Object {
	return receiver.jClass
}

func (receiver *Class) Loader() *ClassLoader {
	return receiver.loader
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
	class.sourceFile = getSourceFile(cf)
	return class
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
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

func (receiver *Class) InitStarted() bool {
	return receiver.initStarted
}

func (receiver *Class) StartInit() {
	receiver.initStarted = true
}

func (receiver *Class) isAccessibleTo(other *Class) bool {
	return receiver.IsPublic() ||
		receiver.GetPackageName() == other.GetPackageName()
}

func (receiver *Class) GetPackageName() string {
	if i := strings.LastIndex(receiver.name, "/"); i >= 0 {
		return receiver.name[:i]
	}
	return ""
}

func (receiver *Class) NewObject() *Object {
	return newObject(receiver)
}

func (receiver *Class) GetMainMethod() *Method {
	return receiver.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (receiver *Class) GetStaticMethod(name, descriptor string) *Method {
	for _, method := range receiver.methods {
		if method.IsStatic() && method.name == name && descriptor == method.descriptor {
			return method
		}
	}
	return nil
}

func (receiver *Class) GetClinitMethod() *Method {
	return receiver.GetStaticMethod("<clinit>", "()V")
}

func (receiver *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(receiver.name)
	return receiver.loader.LoadClass(arrayClassName)
}

func (receiver *Class) isJlObject() bool {
	return receiver.name == "java/lang/Object"
}
func (receiver *Class) isJlCloneable() bool {
	return receiver.name == "java/lang/Cloneable"
}
func (receiver *Class) isJioSerializable() bool {
	return receiver.name == "java/io/Serializable"
}

func (receiver *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[receiver.name]
	return ok
}

func (receiver *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := receiver; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (receiver *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := receiver; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (receiver *Class) JavaName() string {
	return strings.Replace(receiver.name, "/", ".", -1)
}

func (receiver *Class) GetInstanceMethod(name, descriptor string) *Method {
	return receiver.getMethod(name, descriptor, false)
}

func (receiver *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := receiver.getField(fieldName, fieldDescriptor, true)
	return receiver.staticVars.GetRef(field.slotId)
}
func (receiver *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := receiver.getField(fieldName, fieldDescriptor, true)
	receiver.staticVars.SetRef(field.slotId, ref)
}

func (receiver *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(receiver.fields))
		for _, field := range receiver.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return receiver.fields
	}
}
func (receiver *Class) GetConstructor(descriptor string) *Method {
	return receiver.GetInstanceMethod("<init>", descriptor)
}

func (receiver *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(receiver.methods))
	for _, method := range receiver.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}

func (receiver *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(receiver.methods))
	for _, method := range receiver.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}
