package avl

type Visitor interface {
	Visit(node *BstNode)
}

func (node *BstNode) InOrder(visitor Visitor) {
	if node != nil {
		node.Left.InOrder(visitor)
		visitor.Visit(node)
		node.Right.InOrder(visitor)
	}
}

type Counter struct {
	Count int
}

func (counter *Counter) Visit(*BstNode) {
	counter.Count++
}

type Summer struct {
	Sum int
}

func (summer *Summer) Visit(node *BstNode) {
	summer.Sum += node.Value
}
