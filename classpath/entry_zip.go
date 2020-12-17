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

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if zipEntry.zipRC == nil {
		err := zipEntry.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := zipEntry.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, zipEntry, err
}

func (zipEntry *ZipEntry) openJar() error {
	r, err := zip.OpenReader(zipEntry.absPath)
	if err == nil {
		zipEntry.zipRC = r
	}
	return err
}

func (zipEntry *ZipEntry) findClass(className string) *zip.File {
	for _, f := range zipEntry.zipRC.File {
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

func (zipEntry *ZipEntry) String() string {
	return zipEntry.absPath
}
