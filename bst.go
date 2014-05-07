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

func (node *Node) Delete(value int) {
	if node.Right != nil && value > node.Value {
		if node.Right.Value == value {
			node.Right = nil // TODO: obviously broken if Right has children
		}
	} else if node.Left != nil && value < node.Value {
		if node.Left.Value == value {
			node.Left = nil // TODO: obviously broken if Left has children
		}
	}
}

func (node *Node) minimum(parent *Node, depth int) (*Node, *Node, int) {
	if node.Left == nil {
		return node, parent, depth
	}

	return node.Left.minimum(node, depth + 1)
}

func (node *Node) maximum(parent *Node, depth int) (*Node, *Node, int) {
	if node.Right == nil {
		return node, parent, depth
	}

	return node.Right.maximum(node, depth + 1)
}

func (node *Node) Minimum() *Node {
	min, _, _ := node.minimum(nil, 0)
	return min
}

func (node *Node) Maximum() *Node {
	max, _, _ := node.maximum(nil, 0)
	return max
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

func (bst *BinarySearchTree) makeRoot(node *Node) {
	if bst.root.Left != node {
		node.Left = bst.root.Left
	}

	if bst.root.Right != node {
		node.Right = bst.root.Right
	}

	bst.root = node
}

func (bst *BinarySearchTree) Delete(value int) {
	if bst.root != nil {
		if bst.root.Value != value {
			bst.root.Delete(value)
			return
		}

		if bst.root.Left != nil && bst.root.Right != nil {
			maxLeft, parentLeft, leftDepth := bst.root.Left.maximum(bst.root, 0)
			minRight, parentRight, rightDepth := bst.root.Right.minimum(bst.root, 0)

			// Replace the root with the deeper of the max of the left sub-tree and the min of the right sub-tree
			if leftDepth > rightDepth {
				if bst.root != parentLeft {
					parentLeft.Right = nil
				}

				bst.makeRoot(maxLeft)
			} else {
				if bst.root != parentRight {
					parentRight.Left = nil
				}

				bst.makeRoot(minRight)
			}
		} else if bst.root.Left != nil {
			maxLeft, parentLeft, _ := bst.root.Left.maximum(bst.root, 0)

			if bst.root != parentLeft {
				parentLeft.Right = nil
			}

			bst.makeRoot(maxLeft)
		} else if bst.root.Right != nil {
			minRight, parentRight, _ := bst.root.Right.minimum(bst.root, 0)

			if bst.root != parentRight {
				parentRight.Left = nil
			}

			bst.makeRoot(minRight)
		} else {
			// There weren't any children, so back to the empty tree...
			bst.root = nil
		}
	}
}


//        5
//    3       7
//  2   4   6   8
