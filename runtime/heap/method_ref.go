package heap

import (
	"github.com/tinycedar/classp/classfile"
)

type MethodRef struct {
	className  string
	name       string
	descriptor string

	class  *Class
	method *Method
}

func NewMethodRef(refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
	return ref
}

func (m *MethodRef) ResolveMethod() *Method {
	if m.method == nil {
		class := m.ResolveClass()
		m.method = class.FindMethod(m.name, m.descriptor)
	}
	return m.method
}

func (m *MethodRef) ResolveClass() *Class {
	if m.class == nil {
		m.class = LoadClass(m.className)
	}
	return m.class
}
