package main

import (
	"github.com/tinycedar/vanilla/runtime"
	"github.com/tinycedar/vanilla/runtime/heap"
	"github.com/tinycedar/vanilla/runtime/thread"
)

func main() {
	t := thread.NewThread()
	t.Push(thread.NewFrame(t, getMainMethod()))
	runtime.Interpret(t)
}

func getMainMethod() *heap.Method {
	classLoader := runtime.BootstrapClassLoader("./test")
	return classLoader.LoadClass("Test").FindMainMethod()
}
