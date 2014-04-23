package avl

import (
	"testing"
	"time"
	"math/rand"
)

func TestMinimum(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500

	tree := &BinarySearchTree{root}

	expectedMinimum := 1000

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		tree.Insert(value)

		if value < expectedMinimum {
			expectedMinimum = value
		}
	}

	actualMinimum := root.Minimum().Value

	if actualMinimum != expectedMinimum {
		t.Error("expected:",expectedMinimum,", actual:",actualMinimum)
	}
}

func TestMaximum(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500

	tree := &BinarySearchTree{root}

	expectedMaximum := 0

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		tree.Insert(value)

		if value > expectedMaximum {
			expectedMaximum = value
		}
	}

	actualMaximum := root.Maximum().Value

	if actualMaximum != expectedMaximum {
		t.Error("expected:",expectedMaximum,", actual:",actualMaximum)
	}
}

func TestSearchTargetAtRoot(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500
	root.Insert(400)
	root.Insert(600)

	tree := &BinarySearchTree{root}

	foundNode := tree.Search(500)

	if foundNode != root {
		t.Error("expected:",root,", actual:",foundNode)
	}
}

func TestSearchTargetAtLeft(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500
	root.Insert(400)
	root.Insert(600)

	tree := &BinarySearchTree{root}

	foundNode := tree.Search(400)

	if foundNode != root.Left {
		t.Error("expected:",root.Left,", actual:",foundNode)
	}
}

func TestSearchTargetAtRight(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500
	root.Insert(400)
	root.Insert(600)

	tree := &BinarySearchTree{root}

	foundNode := tree.Search(600)

	if foundNode != root.Right {
		t.Error("expected:",root.Right,", actual:",foundNode)
	}
}

func TestDeleteForOneNodeTree(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &Node{}
	root.Value = 500

	tree := &BinarySearchTree{root}

	tree.Delete(500)

	if tree.root != nil {
		t.Error("tree should be empty, root is:",tree.root)
	}
}
