package clrs_test

import (
	"github.com/cpustejovsky/algorithms/clrs"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	got := clrs.MergeSort(input, 4)
	if !EqualIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	got = clrs.MergeSort(input, 4)
	if !EqualIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}
