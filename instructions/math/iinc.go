package math

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type iinc struct {
	opCode   uint8
	index    uint8
	constant int8
}

func (i *iinc) Execute(f *thread.Frame) {
	val := f.LocalVars().GetInt(int(i.index))
	f.LocalVars().SetInt(int(i.index), val+int32(i.constant))
}
