package avl

import (
	"testing"
	"fmt"
)

func TestBreadthFirst(t *testing.T) {
	root := &Node{nil, 500, nil}
	tree := &BinarySearchTree{root}

	tree.Insert(400)
	tree.Insert(300)
	tree.Insert(450)

	tree.Insert(600)
	tree.Insert(700)
	tree.Insert(550)


	lister := &Lister{}

	tree.BreadthFirst(lister)

	expected := []int{500, 400, 600, 300, 450, 550, 700}

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", lister.Elements) {
		t.Error("expected:",expected,", actual:",lister.Elements)
	}
}
