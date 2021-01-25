package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []interface{}
}

func NewStack() *Stack {
	return &Stack{Elements: make([]interface{}, 0)}
}

func (s *Stack) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack) Push(e interface{}) {
	s.Elements = append(s.Elements, e)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("underflow")
	}
	length := len(s.Elements)
	out := s.Elements[length-1]
	s.Elements = s.Elements[:length-1]
	return out, nil
}

func (s *Stack) Print() {
	for _, element := range s.Elements {
		fmt.Printf("%v ", element)
	}
	fmt.Println()
}
