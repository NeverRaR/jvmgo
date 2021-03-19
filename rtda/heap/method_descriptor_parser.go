package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (receiver *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	receiver.raw = descriptor
	receiver.parsed = &MethodDescriptor{}
	receiver.startParams()
	receiver.parseParamTypes()
	receiver.endParams()
	receiver.parseReturnType()
	receiver.finish()
	return receiver.parsed
}

func (receiver *MethodDescriptorParser) startParams() {
	if receiver.readUint8() != '(' {
		receiver.causePanic()
	}
}
func (receiver *MethodDescriptorParser) endParams() {
	if receiver.readUint8() != ')' {
		receiver.causePanic()
	}
}
func (receiver *MethodDescriptorParser) finish() {
	if receiver.offset != len(receiver.raw) {
		receiver.causePanic()
	}
}

func (receiver *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + receiver.raw)
}

func (receiver *MethodDescriptorParser) readUint8() uint8 {
	b := receiver.raw[receiver.offset]
	receiver.offset++
	return b
}
func (receiver *MethodDescriptorParser) unreadUint8() {
	receiver.offset--
}

func (receiver *MethodDescriptorParser) parseParamTypes() {
	for {
		t := receiver.parseFieldType()
		if t != "" {
			receiver.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (receiver *MethodDescriptorParser) parseReturnType() {
	if receiver.readUint8() == 'V' {
		receiver.parsed.returnType = "V"
		return
	}

	receiver.unreadUint8()
	t := receiver.parseFieldType()
	if t != "" {
		receiver.parsed.returnType = t
		return
	}

	receiver.causePanic()
}

func (receiver *MethodDescriptorParser) parseFieldType() string {
	switch receiver.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return receiver.parseObjectType()
	case '[':
		return receiver.parseArrayType()
	default:
		receiver.unreadUint8()
		return ""
	}
}

func (receiver *MethodDescriptorParser) parseObjectType() string {
	unread := receiver.raw[receiver.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		receiver.causePanic()
		return ""
	} else {
		objStart := receiver.offset - 1
		objEnd := receiver.offset + semicolonIndex + 1
		receiver.offset = objEnd
		descriptor := receiver.raw[objStart:objEnd]
		return descriptor
	}
}

func (receiver *MethodDescriptorParser) parseArrayType() string {
	arrStart := receiver.offset - 1
	receiver.parseFieldType()
	arrEnd := receiver.offset
	descriptor := receiver.raw[arrStart:arrEnd]
	return descriptor
}
