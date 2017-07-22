package constants

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type dconst_x struct {
	opCode   uint8
	constVal float64
}

func (l *dconst_x) Execute(f *thread.Frame) {
	f.OperandStack().PushDouble(l.constVal)
}
