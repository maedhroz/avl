package avl

type Node struct {
	Left *Node
	Key int
	Value string
	Right *Node
}

func (*Node) Insert(value int) {
}

func (*Node) Remove(value int) bool {
	return false
}

func (*Node) Find(value int) string {
	return nil
}
