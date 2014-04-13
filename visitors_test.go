package avl

import (
	"testing"
	"time"
	"math/rand"
)

func TestSummer(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &BstNode{}
	root.Value = 500

	expected := 500

	for i := 0; i < 10; i++ {
		value := rand.Int() % 1000
		root.Insert(value)
		expected += value
	}

	summer := &Summer{}
	root.InOrder(summer)

	if summer.Sum != expected {
		t.Error("expected:",expected,", actual:",summer.Sum)
	}
}

func TestCounter(t *testing.T) {
	rand.Seed(time.Now().Unix())

	root := &BstNode{}
	root.Value = 500

	for i := 0; i < 10; i++ {
		root.Insert(rand.Int()%1000)
	}

	counter := &Counter{}
	root.InOrder(counter)

	if counter.Count != 11 {
		t.Error("expected:",11,", actual:",counter.Count)
	}
}
