package main

import (
	"github.com/tinycedar/classp/classfile"
	"github.com/tinycedar/vanilla/runtime"
	"github.com/tinycedar/vanilla/runtime/heap"
	"github.com/tinycedar/vanilla/runtime/thread"
	"io/ioutil"
	"log"
)

func main() {
	t := thread.NewThread()
	t.Push(thread.NewFrame(t, getMainMethod()))
	runtime.Interpret(t)
}

func getMainMethod() *heap.Method {
	bytes, err := ioutil.ReadFile("test/Test.class")
	if err != nil {
		log.Fatal("Error reading class file")
	}
	return heap.FindMainMethod(classfile.Parse(bytes))
}
