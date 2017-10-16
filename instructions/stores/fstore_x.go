package stores

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type fstore_x struct {
	opCode uint8
	index  uint8
}

func (fs *fstore_x) Execute(f *thread.Frame) {
	val := f.OperandStack().PopFloat()
	f.LocalVars().SetFloat(fs.index, val)
}

func (fs *fstore_x) String() string {
	return fmt.Sprintf("{opcode: 0x%x, fstore_%d}", fs.opCode, fs.index)
}
