package comparisons

import "github.com/tinycedar/vanilla/runtime/thread"

// ifeq = 153 (0x99)
// ifne = 154 (0x9a)
// iflt = 155 (0x9b)
// ifge = 156 (0x9c)
// ifgt = 157 (0x9d)
// ifle = 158 (0x9e)
type if_cond struct {
	opCode uint8
	offset int16
}

func (i *if_cond) Execute(f *thread.Frame) {
	succeed := false
	value := f.OperandStack().PopInt()
	switch i.opCode {
	case 0x99:
		succeed = value == 0
	case 0x9a:
		succeed = value != 0
	case 0x9b:
		succeed = value < 0
	case 0x9c:
		succeed = value >= 0
	case 0x9d:
		succeed = value > 0
	case 0x9e:
		succeed = value <= 0
	}
	if succeed {
		f.SetPC(f.GetPC() + int(i.offset))
	}
}
