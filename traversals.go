package avl

func (node *Node) InOrder(visitor Visitor) {
	if node != nil {
		node.Left.InOrder(visitor)
		visitor.Visit(node)
		node.Right.InOrder(visitor)
	}
}

func (bst *BinarySearchTree) InOrder(visitor Visitor) {
	if bst.root != nil {
		bst.root.InOrder(visitor)
	}
}

func (node *Node) BreadthFirst(visitor Visitor) {
	current := []*Node{node}

	for len(current) > 0 {
		frontier := make([]*Node, 0, 2*len(current))

		for i := 0; i < len(current); i++ {
			visitor.Visit(current[i])

			if current[i].Left != nil {
				frontier = append(frontier, current[i].Left)
			}

			if current[i].Right != nil {
				frontier = append(frontier, current[i].Right)
			}
		}

		current = frontier
	}
}

func (bst *BinarySearchTree) BreadthFirst(visitor Visitor) {
	if bst.root != nil {
		bst.root.BreadthFirst(visitor)
	}
}
