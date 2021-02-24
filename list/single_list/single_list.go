package single_list

import (
	"bytes"
	"fmt"
)

type Element struct {
	Value interface{}
	next  *Element
}

func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = nil
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Empty() bool {
	return l.root.next == nil
}

func (l *List) Push(v interface{}) *Element {
	if l.Empty() {
		l.root.next = &Element{Value: v}
		l.len++
		return l.root.next
	}

	current := l.Front()
	for current.Next() != nil {
		current = current.Next()
	}
	current.next = &Element{Value: v}
	l.len++
	return current.next
}

func (l *List) Remove() *Element {
	if l.Empty() {
		return nil
	}
	current := l.Front()
	prev := current
	for current.Next() != nil {
		prev = current
		current = current.Next()
	}
	prev.next = nil
	l.len--
	return current
}

func (l *List) Front() *Element {
	return l.root.next
}

func (l *List) Back() *Element {
	if l.Empty() {
		return nil
	}
	current := l.Front()
	for current.Next() != nil {
		current = current.Next()
	}
	return current
}

func (l *List) String() string {
	var buf bytes.Buffer
	for current := l.Front(); current != nil; current = current.Next() {
		buf.WriteString(fmt.Sprintf("%v ", current.Value))
	}
	return buf.String()
}
