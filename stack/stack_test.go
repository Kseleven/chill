package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	datas := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	stack := NewStack()
	for _, data := range datas {
		stack.Push(data)
	}

	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Print()
}
