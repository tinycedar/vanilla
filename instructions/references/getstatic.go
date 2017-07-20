package references

import "github.com/tinycedar/vanilla/runtime/thread"

type getstatic struct {
	opCode  uint8
	cpIndex uint16
}

func (i *getstatic) Execute(f *thread.Frame) {

}
