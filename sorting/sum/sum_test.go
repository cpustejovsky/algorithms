package sum_test

import (
	"testing"

	"github.com/cpustejovsky/algorithms/sorting/sum"
)

func TestSumSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 8
	want := true
	got := sum.Sort(arr, target)
	if got != want {
		t.Fatalf("wanted:\t%v; got:\t%v", want, got)
	}
}
