package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (receiver *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	receiver.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range receiver.lineNumberTable {
		receiver.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (receiver *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(receiver.lineNumberTable) - 1; i >= 0; i-- {
		entry := receiver.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
