package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (receiver CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range receiver {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (receiver CompositeEntry) String() string {
	strs := make([]string, len(receiver))
	for i, entry := range receiver {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
