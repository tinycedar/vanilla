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
	method := f.Method()
	s := method.Cp.GetConstantInfo(i.index).(*classfile.ConstantMethodrefInfo)
	newFrame := thread.NewFrame(f.Thread(), method.Class.FindMethod(s))
	if argSlotCount := int(newFrame.Method().ArgSlotCount); argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := f.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	f.Thread().Push(newFrame)
}

func (i *invokestatic) String() string {
	return fmt.Sprintf("{opcode: 0x%x, invokestatic}", i.opCode)
}
