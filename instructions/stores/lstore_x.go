package stores

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type lstore_x struct {
	opCode uint8
	index  uint8
}

func (i *lstore_x) Execute(f *thread.Frame) {
	val := f.OperandStack().PopLong()
	f.LocalVars().SetLong(i.index, val)
}

func (i *lstore_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, lstore_%d}", i.opCode, i.index)
}
