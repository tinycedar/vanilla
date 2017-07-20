package heap

import "github.com/tinycedar/classp/classfile"

type Method struct {
	ClassMember
	Cp               classfile.ConstantPool
	Class            *Class
	Name             string
	Descriptor       string
	CodeAttr         *classfile.CodeAttribute
	parsedDescriptor *MethodDescriptor
	ArgSlotCount     uint
}

func newMethod(class *Class, m classfile.MemberInfo) *Method {
	method := &Method{}
	method.Cp = m.ConstantPool()
	method.Class = class
	method.copyMemberInfo(&m)
	method.Name = m.Name()
	method.Descriptor = m.Descriptor()
	method.CodeAttr = m.CodeAttribute()
	method.parsedDescriptor = parseMethodDescriptor(m.Descriptor())
	method.calcArgSlotCount(method.parsedDescriptor.parameterTypes)
	return method
}

func (m *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		m.ArgSlotCount++
		if paramType == "J" || paramType == "D" {
			m.ArgSlotCount++
		}
	}
	if !m.IsStatic() {
		m.ArgSlotCount++ // `this` reference
	}
}
