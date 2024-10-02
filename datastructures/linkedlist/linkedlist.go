package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func New[T comparable](val T) *Node[T] {
	return &Node[T]{
		Value: val,
	}
}

func Search[T comparable](n **Node[T], value T) *Node[T] {
	for cur := n; (*cur) != nil; cur = &(*cur).Next {
		if &(*cur).Value == &value {
			return *cur
		}
	}
	return nil
}

func Walk[T comparable](n **Node[T], arr *[]T) {
	for cur := n; (*cur) != nil; cur = &(*cur).Next {
		*arr = append(*arr, (*cur).Value)
	}
}

func Insert[T comparable](n **Node[T], val T) {
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
	}
	*cur = &Node[T]{
		Value: val,
	}
}

func InsertWithoutPP[T comparable](n *Node[T], val T) {
	new := &Node[T]{
		Value: val,
	}
	if n == nil {
		n = new
	}
	cur := n
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = new
}

func Delete[T comparable](n **Node[T], value T) {
	var prev **Node[T]
	cur := n
	if (*cur) != nil && (*cur).Value == value {
		*n = (*cur).Next
		return
	}
	for ; (*cur) != nil && (*cur).Value != value; cur = &(*cur).Next {
		prev = cur
	}
	if (*cur) != nil && prev != nil {
		(*prev).Next = (*cur).Next
	}

}

func Reverse[T comparable](n *Node[T]) *Node[T] {
	var prev *Node[T]
	cur := n
	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}
