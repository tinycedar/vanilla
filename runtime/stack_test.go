package runtime

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := newStack()
	//stack.Push(NewFrame(nil, 0, 0, nil))
	//stack.Push(NewFrame(nil, 1, 11, nil))
	//stack.Push(NewFrame(nil, 2, 22, nil))
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
