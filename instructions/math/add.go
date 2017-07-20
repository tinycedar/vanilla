package math

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iadd struct {
	opCode uint8
}

func (i *iadd) Execute(f *thread.Frame) {
	stack := f.OperandStack()
	stack.PushInt(stack.PopInt() + stack.PopInt())
}

func (i *iadd) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iadd}", i.opCode)
}
