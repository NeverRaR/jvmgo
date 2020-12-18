package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

func (receiver *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if receiver.zipRC == nil {
		err := receiver.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := receiver.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, receiver, err
}

func (receiver *ZipEntry) openJar() error {
	r, err := zip.OpenReader(receiver.absPath)
	if err == nil {
		receiver.zipRC = r
	}
	return err
}

func (receiver *ZipEntry) findClass(className string) *zip.File {
	for _, f := range receiver.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (receiver *ZipEntry) String() string {
	return receiver.absPath
}
