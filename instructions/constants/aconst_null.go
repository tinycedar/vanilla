package constants

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type aconst_null struct {
	opCode uint8
}

func (i *aconst_null) Execute(f *thread.Frame) {
	f.OperandStack().PushRef(nil)
}
