package instructions

import (
	"github.com/tinycedar/vanilla/runtime/thread"
)

type Instruction interface {
	Execute(frame *thread.Frame)
}
