package runtime

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := newStack()
	stack.Push(newFrame(1, 11))
	stack.Push(newFrame(2, 22))
	stack.Push(newFrame(3, 33))
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
