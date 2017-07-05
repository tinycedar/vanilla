package runtime

import (
	"errors"
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/heap"
	"io/ioutil"
	"log"
	"strings"
)

type ClassLoader struct {
	classpath string
	classMap  map[string]*heap.Class
}

func BootstrapClassLoader(classpath string) *ClassLoader {
	return &ClassLoader{classpath, map[string]*heap.Class{}}
}

//TODO not thread-safe !
func (loader *ClassLoader) LoadClass(className string) *heap.Class {
	if class, ok := loader.classMap[className]; ok {
		return class
	} else {
		bytes, err := readBytes(loader.classpath, className)
		if err != nil {
			log.Fatal("Error reading class file", err)
		}
		class := heap.NewClass(classfile.Parse(bytes))
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
		if bytes, err := ioutil.ReadFile(cp + "/" + className + ".class"); err == nil {
			return bytes, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Class file not found, classpath: %s, className: %s", classpath, className))
}

func resolve() {

}
