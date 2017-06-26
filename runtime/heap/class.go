package heap

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"strings"
)

type Class struct {
	cf      *classfile.ClassFile
	methods []*Method
}

func NewClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.cf = cf
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) Methods() []*Method {
	return c.methods
}

func (c *Class) FindMethod(constMethodRef *classfile.ConstantMethodrefInfo) *Method {
	for _, m := range c.methods {
		if strings.LastIndex(constMethodRef.String(m.Cp), fmt.Sprintf("%s:%s", m.Name, m.Descriptor)) >= 0 {
			return m
		}
	}
	return nil
}

func FindMainMethod(cf *classfile.ClassFile) *Method {
	for _, m := range NewClass(cf).Methods() {
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
