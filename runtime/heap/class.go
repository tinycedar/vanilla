package heap

import (
	"github.com/tinycedar/classp/classfile"
)

type Class struct {
	classLoader *ClassLoader
	cf          *classfile.ClassFile
	methods     []*Method
}

func NewClass(cf *classfile.ClassFile, loader *ClassLoader) *Class {
	class := &Class{}
	class.classLoader = loader
	class.cf = cf
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) Methods() []*Method {
	return c.methods
}

func (c *Class) FindMethod(name, descriptor string) *Method {
	for _, m := range c.methods {
		if m.Name == name && m.Descriptor == descriptor {
			return m
		}
	}
	return nil
}

func (c *Class) FindMainMethod() *Method {
	for _, m := range c.Methods() {
		if m.Name == "main" && m.Descriptor == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func newMethods(class *Class, members []classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(members))
	for i, m := range members {
		methods[i] = newMethod(class, m)
	}
	return methods
}
