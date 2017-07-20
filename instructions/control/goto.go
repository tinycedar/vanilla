package control

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type _goto struct {
	opCode     uint8
	branchoffs int16
}

func (i *_goto) Execute(f *thread.Frame) {
	f.SetPC(f.GetPC() + int(i.branchoffs))
}
