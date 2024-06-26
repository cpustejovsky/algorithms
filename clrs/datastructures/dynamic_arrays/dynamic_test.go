package dynamic_arrays_test

import (
	"github.com/cpustejovsky/algorithms/clrs/datastructures/dynamic_arrays"
	"testing"
)

func TestNew(t *testing.T) {
	size := 8
	da := dynamic_arrays.New(size)
	if da.Cap() != size {
		t.Fatalf("got %v\nwanted %v\n", da.Cap(), size)
	}
	if da.Len() != size {
		t.Fatalf("got %v\nwanted %v\n", da.Len(), size)
	}
}

func TestDynamicArray_Insert(t *testing.T) {
	size := 8
	da := dynamic_arrays.New(size)
	var index int
	if da.Cap() != size {
		t.Fatalf("got %v\nwanted %v\n", da.Cap(), size)
	}
	for i := 1; i <= 10; i++ {
		da.Insert(i)
		if da.Index(index) != i {
			t.Fatalf("got %v\nwanted %v\n", da.Index(index), i)
		}
		index++
	}
	if da.Cap() != size*2 {
		t.Fatalf("got %v\nwanted %v\n", da.Cap(), size*2)
	}
}
