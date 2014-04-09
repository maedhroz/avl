package avl

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func (*Node) Insert(value int) {
}

func (*Node) Remove(value int) bool {
	return false
}
