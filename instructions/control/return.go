package control

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type ireturn struct {
	opCode uint8
}

func (i *ireturn) Execute(f *thread.Frame) {
	val := f.OperandStack().PopInt()
	f.NextFrame().OperandStack().PushInt(val)
	f.Thread().Pop()
}

type _return struct {
	opCode uint8
}

func (i *_return) Execute(f *thread.Frame) {
	f.Thread().Pop()
}

func (i *_return) String() string {
	return fmt.Sprintf("{opcode: 0x%x, return}", i.opCode)
}

// type assertion
func __return() {

}
