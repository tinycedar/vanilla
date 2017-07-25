package loads

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type lload_x struct {
	opCode uint8
	index  uint8
}

func (i *lload_x) Execute(f *thread.Frame) {
	val := f.LocalVars().GetLong(i.index)
	f.OperandStack().PushLong(val)
}

func (i *lload_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, lload_%d}", i.opCode, i.index)
}
