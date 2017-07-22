package constants

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type lconst_x struct {
	opCode   uint8
	constVal int64
}

func (l *lconst_x) Execute(f *thread.Frame) {
	f.OperandStack().PushLong(l.constVal)
}
