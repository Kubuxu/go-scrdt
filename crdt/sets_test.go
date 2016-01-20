//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//

package crdt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ DData = (*GSet)(nil)

func one(i interface{}, err error) interface{} {
	if err == nil {
		return i
	} else {
		return nil
	}
}

func TestGSet(t *testing.T) {
	assert := assert.New(t)

	set1 := NewGSet()
	set1.Add("1")
	set1.Add("2")

	set2 := NewGSet()
	set2.Add("2")
	set2.Add("3")

	assert.True(set1.Contains("1"))
	assert.True(set1.Contains("2"))
	assert.False(set1.Contains("3"))

	assert.False(set2.Contains("1"))
	assert.True(set2.Contains("2"))
	assert.True(set2.Contains("3"))

	set1.MergeIn(set2)

	assert.True(set1.Contains("3"))
	assert.False(set2.Contains("1"))

	assert.Equal([]byte(`["1","2","3"]`), one(set1.MarshalJSON()))

	set3 := NewGSet()
	assert.Nil(set3.UnmarshalJSON([]byte(`["2", "3"]`)))
	assert.Equal(set2, set3)
}

var _ DData = (*TwoPSet)(nil)

func TestTwoPSet(t *testing.T) {
	assert := assert.New(t)
	set1 := NewTwoPSet()

	set1.Add("1")
	set1.Add("2")

	set2 := NewTwoPSet()
	set2.Add("2")
	set2.Add("3")

	assert.True(set1.Contains("1"))
	assert.True(set1.Contains("2"))
	assert.False(set1.Contains("3"))

	assert.False(set2.Contains("1"))
	assert.True(set2.Contains("2"))
	assert.True(set2.Contains("3"))

	set1.MergeIn(set2)

	assert.True(set1.Contains("3"))
	assert.False(set2.Contains("1"))

	set2.Remove("3")
	set1.MergeIn(set2)
	assert.False(set1.Contains("3"))
	assert.True(set1.Removed("3"))

	assert.Equal([]byte(`{"A":["1","2","3"],"R":["3"]}`), one(set1.MarshalJSON()))
}
