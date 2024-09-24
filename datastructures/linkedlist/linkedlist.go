package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func New(val int) *Node {
	return &Node{
		Value: val,
	}
}

func Search(n **Node, value int) *Node {
	for cur := n; (*cur) != nil; cur = &(*cur).Next {
		if &(*cur).Value == &value {
			return *cur
		}
	}
	return nil
}

func Walk(n **Node, arr *[]int) {
	for cur := n; (*cur) != nil; cur = &(*cur).Next {
		*arr = append(*arr, (*cur).Value)
	}
}

func Insert(n **Node, val int) {
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
	}
	*cur = &Node{
		Value: val,
	}
}

func InsertWithoutPP(n *Node, val int) {
	newNode := &Node{
		Value: val,
	}
	if n == nil {
		n = newNode
	}
	cur := n
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
}

func Delete(n **Node, value int) {
	var prev **Node
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

func Reverse(n *Node) {
	a := n.Next
	b := n
	for a != nil {
		tmp := a.Next
		a.Next = b
		b = a
		a = tmp
	}
}