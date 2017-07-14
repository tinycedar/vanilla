package runtime

import (
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime/thread"
	"log"
)

func Interpret(t *thread.Thread) {
	for t.StackNotEmpty() {
		frame := t.CurrentFrame()
		reader := instructions.NewByteCodeReader(frame.Method().Code.Code, frame.NextPC())
		ins := reader.FetchInstruction()
		log.Printf("(%s), %d: %v, \n", frame.Method().Name, frame.NextPC(), ins)
		ins.Execute(frame)
		frame.SetNextPC(reader.NextPC())
	}
}
