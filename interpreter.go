package main

import (
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime"
	"github.com/tinycedar/vanilla/runtime/heap"
	"io/ioutil"
	"log"
)

func bootstrap() {
	thread := runtime.NewThread()
	thread.Push(runtime.NewFrame(thread, getMainMethod()))
	interpret(thread)
}

func getMainMethod() *heap.Method {
	bytes, err := ioutil.ReadFile("test/Test.class")
	if err != nil {
		log.Fatal("Error reading class file")
	}
	class := heap.NewClass(classfile.Parse(bytes))
	for _, m := range class.Methods() {
		if m.Name == "main" && m.Descriptor == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func interpret(t *runtime.Thread) {
	for {
		frame := t.CurrentFrame()
		pc := frame.NextPC()
		byteCodeReader := instructions.NewByteCodeReader(frame.Method().Code.Code, pc)
		ins := byteCodeReader.FetchInstruction()
		//fmt.Printf("pc = %d, ins = %v, methodName: %s, frame: %s\n", pc, ins, frame.Method().Name, frame)
		if ins == nil {
			//TODO to be removed
			pc++
			frame.SetNextPC(pc)
			continue
		}
		ins.Execute(frame)
		frame.SetNextPC(byteCodeReader.NextPC())
		if t.Stack().Size() <= 0 {
			break
		}
		//fmt.Printf("pc = %d, ins = %v, frame: %s\n", pc, ins, frame)
	}
}
