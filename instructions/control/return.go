package control

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type xreturn struct {
	opCode uint8
}

func (i *xreturn) Execute(f *thread.Frame) {
	switch i.opCode {
	case 0xac:
		val := f.OperandStack().PopInt()
		f.NextFrame().OperandStack().PushInt(val)
	case 0xad:
		val := f.OperandStack().PopLong()
		f.NextFrame().OperandStack().PushLong(val)
	case 0xae:
		val := f.OperandStack().PopFloat()
		f.NextFrame().OperandStack().PushFloat(val)
	case 0xaf:
		val := f.OperandStack().PopDouble()
		f.NextFrame().OperandStack().PushDouble(val)
	}
	f.Thread().Pop()
}

//TODO to be removed
type _return struct {
	opCode uint8
}

func (i *_return) Execute(f *thread.Frame) {
	f.Thread().Pop()
}

func (i *_return) String() string {
	return fmt.Sprintf("{opcode: 0x%x, xeturn}", i.opCode)
}
