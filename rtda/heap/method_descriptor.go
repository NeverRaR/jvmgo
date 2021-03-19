package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (receiver *MethodDescriptor) addParameterType(t string) {
	pLen := len(receiver.parameterTypes)
	if pLen == cap(receiver.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, receiver.parameterTypes)
		receiver.parameterTypes = s
	}

	receiver.parameterTypes = append(receiver.parameterTypes, t)
}
