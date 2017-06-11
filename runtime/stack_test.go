package runtime

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := newStack()
	stack.Push(NewFrame(1, 11))
	stack.Push(NewFrame(2, 22))
	stack.Push(NewFrame(3, 33))
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
