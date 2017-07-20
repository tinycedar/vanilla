package comparisons

import "github.com/tinycedar/vanilla/runtime/thread"

// if_acmpeq = 165 (0xa5)
// if_acmpne = 166 (0xa6)

type if_acmp_cond struct {
	opCode uint8
	offset int16
}

func (i *if_acmp_cond) Execute(f *thread.Frame) {
	succeed := false
	value2 := f.OperandStack().PopRef()
	value1 := f.OperandStack().PopRef()
	switch i.opCode {
	case 0xa5:
		succeed = value1 == value2
	case 0xa6:
		succeed = value1 != value2
	}
	if succeed {
		f.SetPC(f.GetPC() + int(i.offset))
	}
}
