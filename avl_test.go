package avl

import (
	"testing"
	"time"
	"math/rand"
)

func TestSummer(t *testing.T) {
	rand.Seed(time.Now().Unix())

	tree := &Node{nil, 500, nil}

	expected := 500

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		tree.Insert(value)
		expected += value
	}

	summer := &Summer{}
	tree.InOrder(summer)

	if summer.Sum != expected {
		t.Error("expected:",expected,", actual:",summer.Sum)
	}
}

func TestCounter(t *testing.T) {
	rand.Seed(time.Now().Unix())

	tree := &Node{nil, 500, nil}

	for i := 0; i < 10; i++ {
		tree.Insert(rand.Int()%1000)
	}

	counter := &Counter{}
	tree.InOrder(counter)

	if counter.Count != 11 {
		t.Error("expected:",11,", actual:",counter.Count)
	}
}
