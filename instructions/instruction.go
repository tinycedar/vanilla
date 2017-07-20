package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type Instruction interface {
	Execute(frame *thread.Frame)
}

type NotImplemented struct {
	opCode uint8
}

func NewNotImplemented(opCode uint8) Instruction {
	return &NotImplemented{opCode}
}

func (i *NotImplemented) Execute(f *thread.Frame) {

}

func (i *NotImplemented) String() string {
	return fmt.Sprintf("{opcode: 0x%x, notImplemented}", i.opCode)
}
