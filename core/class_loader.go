package core

import (
	. "github.com/tinycedar/vanilla/core/class"
	"log"
)

type ClassLoader struct {
	classpath string
	classMap  map[string]*Class
}

func BootstrapClassLoader(classpath string) *ClassLoader {
	return &ClassLoader{classpath, map[string]*Class{}}
}

func (self *ClassLoader) LoadClass(className string) *Class {
	if class, ok := self.classMap[className]; ok {
		return class
	} else {
        // load and create Class by className
		self.classMap[className] = &Class{}
		return self.classMap[className]
	}
}

func resolve() {

}
