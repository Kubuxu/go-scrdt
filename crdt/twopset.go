//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//

package crdt

import (
	"encoding/json"
)

// Strucutre for 2PSet.
type TwoPSet struct {
	add    *GSet
	remove *GSet
}

// Constructs Two Phase Set (2PSet). A set that value can be added and removed only once.
func NewTwoPSet() *TwoPSet {
	return &TwoPSet{NewGSet(), NewGSet()}
}

// Adds value to the set.
func (set *TwoPSet) Add(what string) {
	set.add.Add(what)
}

// Returns length of the set.
func (set *TwoPSet) Remove(what string) {
	set.remove.Add(what)
}

// Returns all elements in set.
func (set *TwoPSet) Len() int {
	return set.add.Len() - set.remove.Len()
}

func (set *TwoPSet) All() []string {
	var result []string
	for _, toAdd := range set.add.All() {
		if !set.remove.Contains(toAdd) {
			result = append(result, toAdd)
		}
	}
	// No need for sorting as GSet is sroted.
	return result
}

// Returns true if element is contained in the set
func (set *TwoPSet) Contains(key string) bool {
	return set.add.Contains(key) && (!set.remove.Contains(key))
}

// Returns true if element was removed from the set
func (set *TwoPSet) Removed(key string) bool {
	return set.remove.Contains(key)
}

func (set *TwoPSet) MergeIn(other DData) error {
	o, ok := other.(*TwoPSet)
	if !ok {
		return ErrInvalidTypeForMerge
	}
	set.add.MergeIn(o.add)
	set.remove.MergeIn(o.remove)
	return nil
}

type rawTwoPSet struct {
	A *GSet
	R *GSet
}

func (gset *TwoPSet) UnmarshalJSON(jsonBlob []byte) error {
	var raw rawTwoPSet
	err := json.Unmarshal(jsonBlob, &raw)
	if err != nil {
		return err
	}
	gset.add.MergeIn(raw.A)
	gset.remove.MergeIn(raw.R)
	return nil
}
func (gset *TwoPSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(rawTwoPSet{gset.add, gset.remove})
}
