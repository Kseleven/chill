package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	datas := []int{1, 2, 3, 5, 3, 7, 8, 10}
	queue := NewQueue()
	for _, data := range datas {
		queue.EnQueue(data)
	}
	queue.Print()
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	queue.Print()
}
