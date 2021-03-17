package heap

import "testing"

func TestHeap(t *testing.T) {
	h := NewHeap()
	h.Insert(1)
	h.Insert(3)
	h.Insert(5)
	h.Insert(2)
	h.Insert(4)
	h.Insert(7)
	h.Insert(9)
	h.RemoveMax()
	t.Log(h.Elements)
}
