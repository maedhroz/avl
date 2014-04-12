package avl

type Node struct {
	Left *Node
	Value int
	Right *Node
}

type Visitor interface {
	Visit(node *Node)
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

func (node *Node) InOrder(visitor Visitor) {
	if node != nil {
		node.Left.InOrder(visitor)
		visitor.Visit(node)
		node.Right.InOrder(visitor)
	}
}

type Counter struct {
	Count int
}

func (counter *Counter) Visit(*Node) {
	counter.Count++
}

type Summer struct {
	Sum int
}

func (summer *Summer) Visit(node *Node) {
	summer.Sum += node.Value
}
