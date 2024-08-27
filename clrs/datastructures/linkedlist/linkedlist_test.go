package linkedlist_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/clrs/datastructures/linkedlist"
)

func TestNew(t *testing.T) {
	want := 42
	got := linkedlist.New(42)
	if want != got.Value {
		t.Fatalf("got %v, wanted %v\n", got.Value, want)
	}
}

func TestInsert(t *testing.T) {
	ll := linkedlist.New(42)

	var arr []int
	linkedlist.Walk(&ll, &arr)

	if len(arr) != 1 {
		t.Fatalf("got %v, wanted %v\n", len(arr), 1)
	}

	linkedlist.Insert(&ll, 43)

	arr = []int{}
	linkedlist.Walk(&ll, &arr)
	if len(arr) != 2 {
		t.Fatalf("got %v, wanted %v\n", len(arr), 2)
	}
}
