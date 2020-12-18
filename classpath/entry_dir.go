package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (receiver *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(receiver.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, receiver, err
}
func (receiver *DirEntry) String() string {
	return receiver.absDir
}
