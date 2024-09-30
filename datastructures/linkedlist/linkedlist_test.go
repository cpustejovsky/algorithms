package linkedlist_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/datastructures/linkedlist"
)

func TestNew(t *testing.T) {
	want := 42
	got := linkedlist.New(42)
	if want != got.Value {
		t.Fatalf("got %v, wanted %v\n", got.Value, want)
	}
}

func TestDelete(t *testing.T) {
	ll := linkedlist.New(42)
	linkedlist.Insert(&ll, 43)
	linkedlist.Insert(&ll, 45)
	linkedlist.Insert(&ll, 47)
	linkedlist.Insert(&ll, 50)
	expectedLength := 5
	var arr []int
	linkedlist.Walk(&ll, &arr)
	if len(arr) != expectedLength {
		t.Fatalf("got %v, wanted %v\n", len(arr), expectedLength)
	}
	linkedlist.Delete(&ll, 42)
	expectedLength--
	arr = []int{}
	linkedlist.Walk(&ll, &arr)
	if len(arr) != expectedLength {
		t.Fatalf("got %v, wanted %v\n", len(arr), expectedLength)
	}
	linkedlist.Delete(&ll, 50)
	expectedLength--
	arr = []int{}
	linkedlist.Walk(&ll, &arr)
	if len(arr) != expectedLength {
		t.Fatalf("got %v, wanted %v\n", len(arr), expectedLength)
	}
	linkedlist.Delete(&ll, 45)
	expectedLength--
	arr = []int{}
	linkedlist.Walk(&ll, &arr)
	if len(arr) != expectedLength {
		t.Fatalf("got %v, wanted %v\n", len(arr), expectedLength)
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
func TestInsertWithoutPP(t *testing.T) {
	ll := linkedlist.New(42)

	var arr []int
	linkedlist.Walk(&ll, &arr)

	if len(arr) != 1 {
		t.Fatalf("got %v, wanted %v\n", len(arr), 1)
	}

	linkedlist.InsertWithoutPP(ll, 43)

	arr = []int{}
	linkedlist.Walk(&ll, &arr)
	if len(arr) != 2 {
		t.Fatalf("got %v, wanted %v\n", len(arr), 2)
	}
}

func TestReverse(t *testing.T) {
	ll := linkedlist.New(42)
	linkedlist.Insert(&ll, 43)
	linkedlist.Insert(&ll, 45)
	linkedlist.Insert(&ll, 47)
	linkedlist.Insert(&ll, 50)
	var arr1 []int
	linkedlist.Walk(&ll, &arr1)
	ll = linkedlist.Reverse(ll)
	var arr2 []int
	linkedlist.Walk(&ll, &arr2)
	for i, val := range arr1 {
		if val != arr2[len(arr2)-1-i] {
			t.Fatalf("%v is not a reversal of %v", arr2, arr1)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	ll := linkedlist.New(42)
	for n := 0; n < b.N; n++ {
		linkedlist.Insert(&ll, n)
	}
}

func BenchmarkInsertWithoutPP(b *testing.B) {
	ll := linkedlist.New(42)
	for n := 0; n < b.N; n++ {
		linkedlist.InsertWithoutPP(ll, n)
	}
}
