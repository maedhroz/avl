package avl

import (
	"testing"
	"time"
	"math/rand"
)

func TestMinimum(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &BstNode{}
	root.Value = 500

	expectedMinimum := 1000

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		root.Insert(value)

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

	root := &BstNode{}
	root.Value = 500

	expectedMaximum := 0

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		root.Insert(value)

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

	root := &BstNode{}
	root.Value = 500
	root.Insert(400)
	root.Insert(600)

	foundNode := root.Search(500)

	if foundNode != root {
		t.Error("expected:",root,", actual:",foundNode)
	}
}
