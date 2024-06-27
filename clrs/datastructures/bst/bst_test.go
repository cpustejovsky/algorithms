package bst_test

import (
	"github.com/cpustejovsky/algorithms/clrs/datastructures/bst"
	"math/rand"
	"sort"
	"testing"
)

// TODO: get to 100% test coverage to make sure you understand how to trigger each case with a test
func TestInsert(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)

	var bstArrr []int
	bst.Walk(tree.Root, &bstArrr)
	if !equalSlices(arr, bstArrr) {
		t.Fatalf("got %v, want %v", bstArrr, arr)
	}

}

func TestMin(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	n := bst.Minimum(tree.Root)
	if n.Value != arr[0] {
		t.Fatalf("got %v, want %v", n.Value, arr[0])
	}
}

func TestMax(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	n := bst.Maximum(tree.Root)
	if n.Value != arr[len(arr)-1] {
		t.Fatalf("got %v, want %v", n.Value, arr[len(arr)-1])
	}
}

func TestSearch(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	found := bst.Search(tree.Root, arr[4])
	if found.Value != arr[4] {
		t.Fatalf("got %v, want %v", found.Value, arr[4])
	}
}

func TestSuccessor(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	found := bst.Search(tree.Root, arr[4])
	s := bst.Successor(found)
	if s.Value != arr[5] {
		t.Fatalf("got %v, want %v", found.Value, arr[5])
	}
}

func TestPredecessor(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	found := bst.Search(tree.Root, arr[4])
	s := bst.Predecessor(found)
	if s.Value != arr[3] {
		t.Fatalf("got %v, want %v", found.Value, arr[3])
	}
}

func TestDelete(t *testing.T) {
	arr := newRandomIntSlice(10)
	tree := newTree(arr)
	sort.Ints(arr)
	found := bst.Search(tree.Root, arr[4])
	bst.Delete(tree, found)
	var bstArrr []int
	bst.Walk(tree.Root, &bstArrr)
	arr = append(arr[:4], arr[5:]...)
	if !equalSlices(arr, bstArrr) {
		t.Fatalf("got %v, want %v", bstArrr, arr)
	}
}

func newRandomIntSlice(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(length)
	}
	return arr
}

func newTree(arr []int) *bst.Tree {
	tree := bst.Tree{
		Root: nil,
	}

	for _, v := range arr {
		node := bst.Node{
			Value:  v,
			Parent: nil,
			Left:   nil,
			Right:  nil,
		}
		bst.Insert(&tree, &node)
	}
	return &tree
}

func equalSlices[C comparable](a, b []C) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
