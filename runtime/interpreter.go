package runtime

import (
	"fmt"
	"github.com/tinycedar/class-parser/classfile"
	"io/ioutil"
	"log"
)

func Bootstrap() {
	thread := newThread()
	codeAttr := getMainMethod().CodeAttribute()
	thread.stack.Push(newFrame(codeAttr.MaxLocals, codeAttr.MaxStack))
	interpret(thread, codeAttr.Code)
}

func getMainMethod() *classfile.MethodInfo {
	bytes, err := ioutil.ReadFile("test/Sample.class")
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

func interpret(t *thread, code []byte) {
	fmt.Println("code: ", code)
	for {

		break
	}
}
