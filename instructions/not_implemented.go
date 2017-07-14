package instructions

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime/thread"
)

type notImplemented struct {
	opCode uint8
}

func (i *notImplemented) Execute(f *thread.Frame) {

}

func (i *notImplemented) String() string {
	return fmt.Sprintf("{opcode: 0x%x, notImplemented}", i.opCode)
}
