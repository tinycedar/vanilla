package comparisons

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

// if_icmpeq = 159 (0x9f)
// if_icmpne = 160 (0xa0)
// if_icmplt = 161 (0xa1)
// if_icmpge = 162 (0xa2)
// if_icmpgt = 163 (0xa3)
// if_icmple = 164 (0xa4)
type if_icmp_cond struct {
	opCode uint8
	offset int16
}

func (i *if_icmp_cond) Execute(f *thread.Frame) {
	succeed := false
	value2 := f.OperandStack().PopInt()
	value1 := f.OperandStack().PopInt()
	switch i.opCode {
	case 0x9f:
		succeed = value1 == value2
	case 0xa0:
		succeed = value1 != value2
	case 0xa1:
		succeed = value1 < value2
	case 0xa2:
		succeed = value1 >= value2
	case 0xa3:
		succeed = value1 > value2
	case 0xa4:
		succeed = value1 <= value2
	}
	if succeed {
		f.SetPC(f.GetPC() + int(i.offset))
	}
}
