package queue_test

import (
	"errors"
	"testing"

	"github.com/cpustejovsky/algorithms/datastructures/queue"
)

func TestEnqueue(t *testing.T) {
	q := queue.New()
	for i := 1; i < 10; i++ {
		q.Enqueue(i)
		if q.Len() != i {
			t.Fatalf("got %v\nwanted %v\n", q.Len(), i)
		}
	}
}

func TestDequeue(t *testing.T) {
	q := queue.New()
	for i := 1; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 1; i < 10; i++ {
		x, err := q.Dequeue()
		if err != nil {
			t.Fatalf("got %v\nwanted %v\n", err, queue.UnderflowError)
		}
		if x != i {
			t.Fatalf("got %v\nwanted %v\n", x, i)
		}
	}
	_, err := q.Dequeue()
	if !errors.Is(err, queue.UnderflowError) {
		t.Fatalf("got %v\nwanted %v\n", err, queue.UnderflowError)
	}
}
