package constants

import "github.com/tinycedar/vanilla/runtime/thread"

type fconst_x struct {
	opCode   uint8
	constVal float32
}

func (l *fconst_x) Execute(f *thread.Frame) {
	f.OperandStack().PushFloat(l.constVal)
}
