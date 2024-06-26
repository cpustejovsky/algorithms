package Ch2_test

import (
	"github.com/cpustejovsky/algorithms/clrs/Ch2"
	"testing"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 8
	want := true
	got := Ch2.Sum(arr, sum)
	if got != want {
		t.Fatalf("wanted:\t%v; got:\t%v", want, got)
	}
}
