package main

import (
	"reflect"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := tn(7, tn(3, nil, nil), tn(15, tn(9, nil, nil), tn(20, nil, nil)))
	it := NewBSTIterator(root)
	got := []int{}
	for it.HasNext() {
		got = append(got, it.Next())
	}
	if !reflect.DeepEqual(got, []int{3, 7, 9, 15, 20}) {
		t.Errorf("iteration = %v, want ascending order", got)
	}
	if it.HasNext() {
		t.Error("exhausted iterator still claims next")
	}

	empty := NewBSTIterator(nil)
	if empty.HasNext() {
		t.Error("empty tree iterator should have nothing")
	}

	// Interleaved HasNext calls must not consume values.
	it2 := NewBSTIterator(tn(2, tn(1, nil, nil), nil))
	it2.HasNext()
	it2.HasNext()
	if v := it2.Next(); v != 1 {
		t.Errorf("Next = %d, want 1", v)
	}
	if v := it2.Next(); v != 2 {
		t.Errorf("Next = %d, want 2", v)
	}
}
