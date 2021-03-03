package classpath

import (
	"os"
)
import "path/filepath"

type Classpath struct {
	bootClasspath Entry

	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (receiver *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	receiver.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	receiver.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (receiver *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		if cp := os.Getenv("CLASS_PATH"); cp != "" {
			cpOption = cp
		} else {
			cpOption = "."
		}
	}
	receiver.userClasspath = newEntry(cpOption)
}

func (receiver *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := receiver.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := receiver.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return receiver.userClasspath.readClass(className)
}

func (receiver *Classpath) String() string {
	return "userClasspath:" + receiver.userClasspath.String()
}
