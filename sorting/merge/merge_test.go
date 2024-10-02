package merge_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/sorting/merge"
)

func TestSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	got := merge.Sort(input, 4)
	if !equalIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	got = merge.Sort(input, 4)
	if !equalIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}
