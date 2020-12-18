package classfile

type MarkerAttribute struct {
}

func (receiver *MarkerAttribute) readInfo(reader *ClassReader) { // read nothing
}

type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}
