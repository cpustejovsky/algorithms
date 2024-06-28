package bst

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return &Tree{}
}

func Root(t *Tree) *Node {
	return t.root
}

func Search(n *Node, value int) *Node {
	if n == nil || value == n.Value {
		return n
	}
	if value < n.Value {
		return Search(n.Left, value)
	} else {
		return Search(n.Right, value)
	}
}

func Walk(x *Node, arr *[]int) {
	if x != nil {
		Walk(x.Left, arr)
		*arr = append(*arr, x.Value)
		Walk(x.Right, arr)
	}
}

func Minimum(x *Node) *Node {
	if x.Left == nil {
		return x
	}
	return Minimum(x.Left)
}

func Maximum(x *Node) *Node {
	if x.Right == nil {
		return x
	}
	return Maximum(x.Right)
}

// Successor takes a Node and find the next node based on in-order traversal
func Successor(x *Node) *Node {
	if x.Right != nil {
		return Minimum(x.Right)
	}
	parent := x.Parent
	if parent == nil || x != parent.Right {
		return parent
	}
	return Successor(parent)
}

// Predecessor takes a Node and find the preceding node based on in-order traversal
func Predecessor(x *Node) *Node {
	if x.Left != nil {
		return Maximum(x.Right)
	}
	parent := x.Parent
	if parent == nil || x != parent.Left {
		return parent
	}
	return Predecessor(parent)
}

func Insert(t *Tree, new *Node) {
	var p *Node = nil
	x := t.root
	//go down the tree until x is nil
	for x != nil {
		//keep track of parent as p
		p = x
		if new.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	//you've gone too far in getting x to nil
	//now use the parent to insert correctly
	new.Parent = p
	if p == nil {
		t.root = new //in this case, the tree was empty
	} else if new.Value < p.Value {
		p.Left = new
	} else {
		p.Right = new
	}
}

func transplant(t *Tree, u, v *Node) {
	if u.Parent == nil {
		t.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func Delete(t *Tree, z *Node) {
	if z.Left == nil {
		transplant(t, z, z.Right)
	} else if z.Right == nil {
		transplant(t, z, z.Left)
	} else {
		y := Minimum(z.Right)
		if y.Parent != z {
			transplant(t, y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		transplant(t, z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}
