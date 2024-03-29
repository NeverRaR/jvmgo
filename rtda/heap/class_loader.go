package heap

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (receiver *ClassLoader) loadBasicClasses() {
	jlClassClass := receiver.LoadClass("java/lang/Class")
	for _, class := range receiver.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (receiver *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		receiver.loadPrimitiveClass(primitiveType)
	}
}

func (receiver *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      receiver,
		initStarted: true,
	}
	class.jClass = receiver.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	receiver.classMap[className] = class
}

func (receiver *ClassLoader) LoadClass(name string) *Class {
	if class, ok := receiver.classMap[name]; ok {
		return class
	}
	var class *Class
	if name[0] == '[' {
		class = receiver.loadArrayClass(name)
	} else {
		class = receiver.loadNonArrayClass(name)
	}
	if jlClass, ok := receiver.classMap["java/lang/Class"]; ok {
		class.jClass = jlClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

func (receiver *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := receiver.readClass(name)
	class := receiver.defineClass(data)
	link(class)
	if receiver.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (receiver *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //todo
		name:        name,
		loader:      receiver,
		initStarted: true,
		superClass:  receiver.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			receiver.LoadClass("java/lang/Cloneable"),
			receiver.LoadClass("java/io/Serializable"),
		},
	}
	receiver.classMap[name] = class
	return class
}

func (receiver *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := receiver.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (receiver *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = receiver
	resolveSuperClass(class)
	resolveInterfaces(class)
	receiver.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.loader, goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
