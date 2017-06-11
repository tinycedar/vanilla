package instructions

import "github.com/tinycedar/vanilla/runtime"

type Instruction interface {
	Execute(frame *runtime.Frame)
}
