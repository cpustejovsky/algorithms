package stack_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cpustejovsky/algorithms/datastructures/stack"
)

func TestStack_Push(t *testing.T) {
	s := stack.New()
	arr := []rune{'1', '2', '3', '4', '5'}
	for _, char := range arr {
		s.Push(char)
	}
	if len(arr) != s.Len() {
		t.Fatalf("got %#v\nwanted %#v\n", s.Len(), len(arr))
	}
}

func TestStack_Pop(t *testing.T) {
	s := stack.New()
	for _, char := range []rune{'1', '2', '3', '4', '5'} {
		s.Push(char)
	}
	wantStack := stack.New()
	for _, char := range []rune{'1', '2', '3', '4'} {
		wantStack.Push(char)
	}
	val, err := s.Pop()
	if err != nil {
		t.Fatalf("got %#v\nwanted %#v\n", err, nil)
	}
	if !reflect.DeepEqual(s, wantStack) {
		t.Fatalf("got %#v\nwanted %#v\n", s, wantStack)
	}
	wantVal := '5'
	if val != wantVal {
		t.Fatalf("got %#v\nwanted %#v\n", val, wantVal)
	}
	empty := stack.New()
	_, err = empty.Pop()
	if !errors.Is(err, stack.UnderflowError) {
		t.Fatalf("got %#v\nwanted %#v\n", err, stack.UnderflowError)
	}
}
