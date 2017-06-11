package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime"
)

type iadd struct {
	opCode uint8
}

func (i *iadd) Execute(f *runtime.Frame) {
	stack := f.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (i *iadd) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iadd}", i.opCode)
}
