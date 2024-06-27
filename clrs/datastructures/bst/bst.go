package bst

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func Search(n *Node, value int) *Node {
	for n != nil && value != n.Value {
		if value < n.Value {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return n
}

func Walk(x *Node, arr *[]int) {
	if x != nil {
		Walk(x.Left, arr)
		*arr = append(*arr, x.Value)
		Walk(x.Right, arr)
	}
}

func Minimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func Maximum(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

// Successor takes a Node and find the next node based on in-order traversal
func Successor(x *Node) *Node {
	if x.Right != nil {
		return Minimum(x.Right)
	}
	parent := x.Parent
	for parent != nil && x == parent.Right {
		x = parent
		parent = parent.Parent
	}
	return parent
}

// Predecessor takes a Node and find the preceding node based on in-order traversal
func Predecessor(x *Node) *Node {
	if x.Left != nil {
		return Maximum(x.Left)
	}
	parent := x.Parent
	for parent != nil && x == parent.Left {
		x = parent
		parent = parent.Parent
	}
	return parent
}

func Insert(t *Tree, new *Node) {
	var p *Node = nil
	x := t.Root
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
		t.Root = new //in this case, the tree was empty
	} else if new.Value < p.Value {
		p.Left = new
	} else {
		p.Right = new
	}
}

func Transplant(t *Tree, u, v *Node) {
	if u.Parent == nil {
		t.Root = v
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
		Transplant(t, z, z.Right)
	} else if z.Right == nil {
		Transplant(t, z, z.Left)
	} else {
		y := Minimum(z.Right)
		if y.Parent != z {
			Transplant(t, y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		Transplant(t, z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}
