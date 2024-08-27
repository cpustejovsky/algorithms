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

func Search(n *Node, value int) *Node {
	if n == nil || value == n.Value {
		return n
	}
	return Search(n.Next, value)
}

func Walk(n **Node, arr *[]int) {
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
		*arr = append(*arr, (*cur).Value)
	}
	// if x != nil {
	// 	Walk(x.Next, arr)
	// 	*arr = append(*arr, x.Value)
	// }
}

func Insert(n **Node, val int) {
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
	}
	*cur = &Node{
		Value: val,
	}
}
