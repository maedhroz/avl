package avl

type Visitor interface {
	Visit(node *Node)
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

type Lister struct {
	Elements []int
}

func (lister *Lister) Visit(node *Node) {
	lister.Elements = append(lister.Elements, node.Value)
}
