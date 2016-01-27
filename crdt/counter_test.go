//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//
package crdt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ DData = (*LCounter)(nil)

func TestLCounter(t *testing.T) {
	assert := assert.New(t)

	l1 := NewLCounter()
	l2 := NewLCounter()

	assert.Equal(0, l1.Value(), "After creation value of counter should be 0.")
	assert.Equal(0, l2.Value(), "After creation value of counter should be 0.")

	l1.Increment()
	assert.Equal(1, l1.Value(), "After increnent value should be one more.")
	l1.Increment()
	l1.Increment()

	l2.Increment()

	l2.MergeIn(l1)

	assert.Equal(3, l2.Value(), "After merge value in counter should be eqal to maximum value.")

	assert.Equal([]byte(`3`), one(l2.MarshalJSON()), "JSON representation should be just a number.")
}

var _ DData = (*GCounter)(nil)

func TestGCounter(t *testing.T) {
	assert := assert.New(t)
	c1 := NewGCounter(NodeID("1"))
	c2 := NewGCounter(NodeID("2"))

	assert.Equal(0, c1.Value(), "After creation value of counter should be 0.")
	assert.Equal(0, c2.Value(), "After creation value of counter should be 0.")

	c1.Increment()
	c1.Increment()
	assert.Equal(2, c1.Value(), "Two increments should make counter to be 2.")

	c2.Increment()
	assert.Equal(1, c2.Value(), "An increment should make counter be equal 1.")

	c3 := NewGCounter(NodeID("3"))
	c3.MergeIn(c2)
	assert.Equal(c2.Value(), c3.Value(), "After merge to empy counter values should be equal.")

	c3.MergeIn(c1)
	assert.Equal(c2.Value()+c1.Value(), c3.Value(), "After merging two counter that haven't seen each other before value should be the sum.")

	c1.MergeIn(c3)
	c2.MergeIn(c3)
	assert.Equal(c3.Value(), c1.Value(), "After propagation all counters should be equal.")
	assert.Equal(c3.Value(), c2.Value(), "After propagation all counters should be equal.")
}
