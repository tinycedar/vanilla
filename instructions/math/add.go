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
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (i *iadd) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iadd}", i.opCode)
}
