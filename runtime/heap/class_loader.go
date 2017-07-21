package heap

import (
	"errors"
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//TODO singleton & java_home/JDK classpath auto scan
//TODO need other system class loaders

var separator = string(os.PathSeparator)

var loader *ClassLoader

type ClassLoader struct {
	classpath string
	classMap  map[string]*Class
}

func InitClassLoader(classpath string) {
	if loader == nil {
		if classpath == "" {
			classpath = "."
		}
		loader = &ClassLoader{classpath, map[string]*Class{}}
	}
}

//TODO not thread-safe !
func LoadClass(className string) *Class {
	if loader == nil {
		panic("ClassLoader is not initialized !")
	}
	if class, ok := loader.classMap[className]; ok {
		return class
	} else {
		bytes, err := readBytes(loader.classpath, className)
		if err != nil {
			log.Fatal("Error reading class file", err)
		}
		class := NewClass(classfile.Parse(bytes), loader)
		loader.classMap[className] = class
		return class
	}
}

//TODO read jar/war etc
func readBytes(classpath, className string) ([]byte, error) {
	for _, cp := range strings.Split(classpath, ":") {
		cp = strings.TrimSpace(cp)
		if cp == "" {
			continue
		}
		path := cp + separator + strings.Replace(className, ".", separator, -1) + ".class"
		if bytes, err := ioutil.ReadFile(path); err == nil {
			return bytes, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Class file not found, classpath: %s, className: %s", classpath, className))
}
