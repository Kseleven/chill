package single_list

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New()
	for i := 0; i < 10; i++ {
		l.Push(i)
	}

	t.Log(l)
	t.Logf("remove:%+v\n", l.Remove())
	t.Logf("remove:%+v\n", l.Remove())
	t.Logf("list:%+v front:%+v back:%+v\n", l, l.Front().Value, l.Back().Value)
}
