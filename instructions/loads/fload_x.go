package loads

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type fload_x struct {
	opCode uint8
	index  uint8
}

func (fl *fload_x) Execute(f *thread.Frame) {
	val := f.LocalVars().GetFloat(fl.index)
	f.OperandStack().PushFloat(val)
}

func (fl *fload_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, fload_%d}", fl.opCode, fl.index)
}
