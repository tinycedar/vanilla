package comparisons

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type if_icmpge struct {
	opCode uint8
	offset uint16
}

func (i *if_icmpge) Execute(f *thread.Frame) {
	value2 := f.OperandStack().PopInt()
	value1 := f.OperandStack().PopInt()
	if value1 >= value2 {
		f.SetPC(f.GetPC() + int(i.offset))
	}
}
