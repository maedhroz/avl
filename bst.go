package avl

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func (node *Node) Insert(value int) {
	if value >= node.Value {
		if node.Right != nil {
			node.Right.Insert(value)
		} else {
			node.Right = &Node{nil, value, nil}
		}
	} else {
		if node.Left != nil {
			node.Left.Insert(value)
		} else {
			node.Left = &Node{nil, value, nil}
		}
	}
}

func (node *Node) Search(value int) *Node {
	if node.Right != nil && value > node.Value {
		return node.Right.Search(value)
	} else if node.Left != nil && value < node.Value{
		return node.Left.Search(value)
	}

	if node.Value == value {
		return node
	}

	return nil
}

func (node *Node) Minimum() *Node {
	if node.Left == nil {
		return node
	}

	return node.Left.Minimum()
}

func (node *Node) Maximum() *Node {
	if node.Right == nil {
		return node
	}

	return node.Right.Maximum()
}


type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(value int) {
	if bst.root == nil {
		bst.root = &Node{nil, value, nil}
	} else {
		bst.root.Insert(value)
	}
}

func (bst *BinarySearchTree) Search(value int) *Node {
	if bst.root == nil {
		return nil
	} else {
		return bst.root.Search(value)
	}
}

func (bst *BinarySearchTree) Delete(value int) {
	if bst.root != nil {
		if bst.root.Value == value {
			// TODO: Broken unless there is only a root
			bst.root = nil
		}
	}
}
