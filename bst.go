package avl

type BstNode struct {
	Left *BstNode
	Value int
	Right *BstNode
}

func (node *BstNode) Insert(value int) {
	if value >= node.Value {
		if node.Right != nil {
			node.Right.Insert(value)
		} else {
			node.Right = &BstNode{nil, value, nil}
		}
	} else {
		if node.Left != nil {
			node.Left.Insert(value)
		} else {
			node.Left = &BstNode{nil, value, nil}
		}
	}
}

func (node *BstNode) Search(value int) *BstNode {
	if value > node.Value && node.Right != nil {
		node.Right.Search(value)
	} else if value < node.Value && node.Left != nil {
		node.Left.Search(value)
	}

	if node.Value == value {
		return node
	}

	return nil
}

func (node *BstNode) Minimum() *BstNode {
	if node.Left == nil {
		return node
	}

	return node.Left.Minimum()
}

func (node *BstNode) Maximum() *BstNode {
	if node.Right == nil {
		return node
	}

	return node.Right.Maximum()
}

