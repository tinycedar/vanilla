package heap

import "github.com/tinycedar/classp/classfile"

type Method struct {
	Cp         classfile.ConstantPool
	Class      *Class
	Name       string
	Descriptor string
	Code       *classfile.CodeAttribute
}

func newMethod(class *Class, m classfile.MemberInfo) *Method {
	return &Method{m.ConstantPool(), class, m.Name(), m.Descriptor(), m.CodeAttribute()}
}
