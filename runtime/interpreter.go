package runtime

import (
	"fmt"
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/instructions/comparisons"
	"github.com/tinycedar/vanilla/instructions/constants"
	"github.com/tinycedar/vanilla/instructions/control"
	"github.com/tinycedar/vanilla/instructions/conversions"
	"github.com/tinycedar/vanilla/instructions/extended"
	"github.com/tinycedar/vanilla/instructions/loads"
	"github.com/tinycedar/vanilla/instructions/math"
	"github.com/tinycedar/vanilla/instructions/references"
	"github.com/tinycedar/vanilla/instructions/stack"
	"github.com/tinycedar/vanilla/instructions/stores"
	"github.com/tinycedar/vanilla/runtime/thread"
)

func Interpret(t *thread.Thread) {
	for t.StackNotEmpty() {
		frame := t.CurrentFrame()
		pc := frame.GetPC()
		reader := instructions.NewByteCodeReader(frame.Method().CodeAttr.Code, pc)
		ins := fetchInstruction(reader)
		//log.Printf("(%s), %d: %v, \n", frame.Method().Name, pc, ins)
		ins.Execute(frame)
		if frame.GetPC() == pc {
			frame.SetPC(reader.Offset())
		}
	}
}

func fetchInstruction(reader *instructions.CodeReader) instructions.Instruction {
	opCode := reader.ReadUint8()
	switch {
	case opCode <= 0x14:
		return constants.FetchInstruction(opCode, reader)
	case opCode <= 0x35:
		return loads.FetchInstruction(opCode, reader)
	case opCode <= 0x56:
		return stores.FetchInstruction(opCode, reader)
	case opCode <= 0x5f:
		return stack.FetchInstruction(opCode, reader)
	case opCode <= 0x84:
		return math.FetchInstruction(opCode, reader)
	case opCode <= 0x93:
		return conversions.FetchInstruction(opCode, reader)
	case opCode <= 0xa6:
		return comparisons.FetchInstruction(opCode, reader)
	case opCode <= 0xb1:
		return control.FetchInstruction(opCode, reader)
	case opCode <= 0xc3:
		return references.FetchInstruction(opCode, reader)
	case opCode <= 0xc9:
		return extended.FetchInstruction(opCode, reader)
	case opCode <= 0xff:
		return instructions.NewNotImplemented(opCode)
	default:
		panic(fmt.Sprintf("invalid opCode: %d", opCode))
	}
}
