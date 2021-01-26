package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Empty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) EnQueue(value interface{}) {
	q.Elements = append(q.Elements, value)
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("underflow")
	}
	out := q.Elements[0]
	q.Elements = q.Elements[1:]
	return out, nil
}

func (q *Queue) Print() {
	for _, element := range q.Elements {
		fmt.Printf("%v ", element)
	}
	fmt.Println()
}
