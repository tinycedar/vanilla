package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type _return struct {
	opCode uint8
}

func (i *_return) Execute(f *thread.Frame) {
	f.Thread().Pop()
}

func (i *_return) String() string {
	return fmt.Sprintf("{opcode: 0x%x, return}", i.opCode)
}
