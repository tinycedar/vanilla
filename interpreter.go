package main

import (
	"github.com/tinycedar/class-parser/classfile"
	"github.com/tinycedar/vanilla/instructions"
	"github.com/tinycedar/vanilla/runtime"
	"io/ioutil"
	"log"
)

func bootstrap() {
	thread := runtime.NewThread()
	method := getMainMethod()
	codeAttr := method.CodeAttribute()
	thread.Stack().Push(runtime.NewFrame(thread, uint(codeAttr.MaxLocals), uint(codeAttr.MaxStack), method))
	interpret(thread, codeAttr.Code)
}

func getMainMethod() *classfile.MethodInfo {
	bytes, err := ioutil.ReadFile("test/Test.class")
	if err != nil {
		log.Fatal("Error reading class file")
	}
	cf := classfile.NewClassFile()
	cf.Read(classfile.NewClassReader(bytes))
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return &m
		}
	}
	return nil
}

func interpret(t *runtime.Thread, code []byte) {
	//fmt.Println("code: ", code)
	frame := t.Stack().Pop()
	byteCodeReader := instructions.NewByteCodeReader(code)
	for {
		pc := t.GetPC()
		if pc >= len(code) {
			break
		}
		ins := byteCodeReader.FetchInstruction()
		//fmt.Printf("pc = %d, ins = %v, frame: %s\n", pc, ins, frame)
		if ins == nil {
			//TODO to be removed
			pc++
			t.SetPC(pc)
			continue
		}
		ins.Execute(frame)
		t.SetPC(byteCodeReader.NextPC())

		//fmt.Printf("pc = %d, ins = %v, frame: %s\n", pc, ins, frame)
	}
}
