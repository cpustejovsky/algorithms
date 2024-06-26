package queue

import "errors"

var UnderflowError error = errors.New("stack underflow")

// Queue data structure has a FIFO policy (first-in, first-out)
type Queue struct {
	array []any
}

func New() Queue {
	return Queue{
		array: make([]any, 0),
	}
}

func (q *Queue) Len() int {
	return len(q.array)
}

func (q *Queue) Enqueue(val any) {
	q.array = append(q.array, val)
}

func (q *Queue) Dequeue() (any, error) {
	if q.Len() < 1 {
		return 0, UnderflowError
	}
	x := q.array[0]
	q.array = q.array[1:]
	return x, nil
}
