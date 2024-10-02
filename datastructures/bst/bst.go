package bst

import "cmp"

type Node[T cmp.Ordered] struct {
	Value  T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

type Tree[T cmp.Ordered] struct {
	root *Node[T]
}

func New[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func Root[T cmp.Ordered](t *Tree[T]) *Node[T] {
	return t.root
}

func Search[T cmp.Ordered](n *Node[T], value T) *Node[T] {
	if n == nil || value == n.Value {
		return n
	}
	if value < n.Value {
		return Search(n.Left, value)
	} else {
		return Search(n.Right, value)
	}
}

func Walk[T cmp.Ordered](x *Node[T], arr *[]T) {
	if x != nil {
		Walk(x.Left, arr)
		*arr = append(*arr, x.Value)
		Walk(x.Right, arr)
	}
}

func Minimum[T cmp.Ordered](x *Node[T]) *Node[T] {
	if x.Left == nil {
		return x
	}
	return Minimum(x.Left)
}

func Maximum[T cmp.Ordered](x *Node[T]) *Node[T] {
	if x.Right == nil {
		return x
	}
	return Maximum(x.Right)
}

// Successor takes a Node[T] and find the next node based on in-order traversal
func Successor[T cmp.Ordered](x *Node[T]) *Node[T] {
	if x.Right != nil {
		return Minimum(x.Right)
	}
	parent := x.Parent
	if parent == nil || x != parent.Right {
		return parent
	}
	return Successor(parent)
}

// Predecessor takes a Node[T] and find the preceding node based on in-order traversal
func Predecessor[T cmp.Ordered](x *Node[T]) *Node[T] {
	if x.Left != nil {
		return Maximum(x.Right)
	}
	parent := x.Parent
	if parent == nil || x != parent.Left {
		return parent
	}
	return Predecessor(parent)
}

func Insert[T cmp.Ordered](t *Tree[T], new *Node[T]) {
	var p *Node[T] = nil
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

func transplant[T cmp.Ordered](t *Tree[T], u, v *Node[T]) {
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

func Delete[T cmp.Ordered](t *Tree[T], z *Node[T]) {
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
