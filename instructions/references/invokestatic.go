package references

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type invokestatic struct {
	opCode uint8
	index  uint16
}

func (i *invokestatic) Execute(f *thread.Frame) {
	s := f.Method().Cp.GetConstantInfo(i.index).(*classfile.ConstantMethodrefInfo)
	f.Thread().Push(thread.NewFrame(f.Thread(), f.Method().Class.FindMethod(s)))
	//TODO add method arg related
}

func (i *invokestatic) String() string {
	return fmt.Sprintf("{opcode: 0x%x, invokestatic}", i.opCode)
}
