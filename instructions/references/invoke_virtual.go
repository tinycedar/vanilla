package references

import (
	"fmt"
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime/thread"
	"strings"
)

type invokevirtual struct {
	opCode uint8
	offset uint16
}

func (i *invokevirtual) Execute(f *thread.Frame) {
	method := f.Method()
	if cp, ok := method.Cp[i.offset].(*classfile.ConstantMethodrefInfo); ok {
		invoked := cp.String()
		if strings.Index(invoked, "java/io/PrintStream.println") >= 0 {
			if strings.LastIndex(invoked, "(Ljava/lang/String;)V") >= 0 {
				fmt.Println(f.OperandStack().PopString())
			} else if strings.LastIndex(invoked, "(I)V") >= 0 {
				fmt.Println(f.OperandStack().PopInt())
			} else if strings.LastIndex(invoked, "(F)V") >= 0 {
				fmt.Println(f.OperandStack().PopFloat())
			} else {
				panic("Not implemented return value: " + invoked)
			}
		}
	}
}

func (i *invokevirtual) String() string {
	return fmt.Sprintf("{opcode: 0x%x, invokevirtual}", i.opCode)
}
