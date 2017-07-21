package main

import (
	"fmt"
	"github.com/tinycedar/vanilla/runtime"
	"github.com/tinycedar/vanilla/runtime/heap"
	"github.com/tinycedar/vanilla/runtime/thread"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Vanilla v0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		heap.InitClassLoader(cmd.classpath)

		thread := thread.NewThread()
		thread.NewFrame(heap.LoadClass(cmd.class).FindMainMethod())
		runtime.Interpret(thread)
	}
}
