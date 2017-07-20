package loads

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iload_x struct {
	opCode uint8
	index  uint8
}

func (i *iload_x) Execute(f *thread.Frame) {
	val := f.LocalVars().GetInt(i.index)
	f.OperandStack().PushInt(val)
}

func (i *iload_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, iload_%d}", i.opCode, i.index)
}
