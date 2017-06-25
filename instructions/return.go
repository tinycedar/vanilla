package instructions

import "github.com/tinycedar/vanilla/runtime"

type _return struct {
	opCode uint8
}

func (i *_return) Execute(f *runtime.Frame) {
	f.Thread().Stack().Pop()
}
