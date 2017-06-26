package runtime

import (
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime/thread"
)

func Interpret(t *thread.Thread) {
	for t.StackNotEmpty() {
		frame := t.CurrentFrame()
		reader := instructions.NewByteCodeReader(frame.Method().Code.Code, frame.NextPC())
		ins := reader.FetchInstruction()
		//fmt.Printf("pc: %d, %v, { %s }, frame: %s\n", frame.NextPC(), ins, frame.Method().Name, frame)
		if ins != nil {
			ins.Execute(frame)
		}
		frame.SetNextPC(reader.NextPC())
	}
}
