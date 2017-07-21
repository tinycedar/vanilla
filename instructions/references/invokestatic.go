package references

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/heap"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type invokestatic struct {
	opCode uint8
	index  uint16
}

func (i *invokestatic) Execute(f *thread.Frame) {
	methodRefInfo := f.Method().Cp.GetConstantInfo(i.index).(*classfile.ConstantMethodrefInfo)
	newFrame := f.Thread().NewFrame(heap.NewMethodRef(methodRefInfo).ResolveMethod())

	if argSlotCount := int(newFrame.Method().ArgSlotCount); argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := f.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}

func (i *invokestatic) String() string {
	return fmt.Sprintf("{opcode: 0x%x, invokestatic}", i.opCode)
}
