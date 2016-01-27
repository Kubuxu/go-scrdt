//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//

package crdt

import (
	"encoding/json"
)

// Grow only counter that is true CRDT.
type GCounter struct {
	id     NodeID
	counts map[NodeID]*LCounter
}

func NewGCounter(nodeid NodeID) *GCounter {
	return &GCounter{nodeid, make(map[NodeID]*LCounter)}
}

func (g *GCounter) Increment() {
	l, ok := g.counts[g.id]
	if !ok {
		l = NewLCounter()
		g.counts[g.id] = l
	}
	l.Increment()
}

func (g *GCounter) Value() int {
	res := 0
	for _, v := range g.counts {
		res += v.Value()
	}
	return res
}

func (g *GCounter) MergeIn(other DData) error {
	o, ok := other.(*GCounter)
	if !ok {
		return ErrInvalidTypeForMerge
	}
	for k, v := range o.counts {
		count, ok := g.counts[k]
		if !ok {
			count = NewLCounter()
			g.counts[k] = count
		}
		err := count.MergeIn(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GCounter) UnmarshalJSON(jsonBlob []byte) error {
	return json.Unmarshal(jsonBlob, &g.counts)
}
func (g *GCounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.counts)
}
