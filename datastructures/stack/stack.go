package stack

import "errors"

// Stack has a LIFO policy of last-in, first-out
type Stack struct {
	array []any
}

var UnderflowError error = errors.New("stack underflow")

func New() *Stack {
	return &Stack{
		array: make([]any, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.array)
}

func (s *Stack) Push(val any) {
	s.array = append(s.array, val)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() (any, error) {
	if s.Empty() {
		return -1, UnderflowError
	}
	return s.array[s.Len()-1], nil
}

func (s *Stack) Pop() (any, error) {
	top, err := s.Top()
	if err != nil {
		return -1, err
	}
	s.array = s.array[:s.Len()-1]
	return top, nil
}
