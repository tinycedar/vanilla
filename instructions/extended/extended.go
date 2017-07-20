package extended

import "github.com/tinycedar/vanilla/instructions"

func FetchInstruction(opCode uint8, reader *instructions.CodeReader) instructions.Instruction {
	switch opCode {
	default:
		return instructions.NewNotImplemented(opCode)
	}
}
