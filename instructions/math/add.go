package math

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type xadd struct {
	opCode uint8
}

func (i *xadd) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	switch i.opCode {
	case 0x60:
		stack.PushInt(stack.PopInt() + stack.PopInt())
	case 0x61:
		stack.PushLong(stack.PopLong() + stack.PopLong())
	case 0x62:
		stack.PushFloat(stack.PopFloat() + stack.PopFloat())
	case 0x63:
		stack.PushDouble(stack.PopDouble() + stack.PopDouble())
	}
}

func (i *xadd) String() string {
	return fmt.Sprintf("{opcode: 0x%x, xadd}", i.opCode)
}
