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
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
		if &(*cur).Value == &value {
			return *cur
		}
	}
	return nil
}

func Walk(n **Node, arr *[]int) {
	cur := n
	for ; (*cur) != nil; cur = &(*cur).Next {
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
