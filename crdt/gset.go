//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//

package crdt

import (
	"encoding/json"
	"sort"
)

// Structure for GSet
type GSet struct {
	set map[string]struct{}
}

// Constructs Grow Only Set (GSet). A set to which values can only be added.
func NewGSet() *GSet {
	return &GSet{make(map[string]struct{})}
}

// Adds value to the set.
func (gset *GSet) Add(value string) {
	gset.set[value] = struct{}{}
}

// Returns length of the set.
func (gset *GSet) Len() int {
	return len(gset.set)
}

//Returns all elements in set.
func (gset *GSet) All() []string {
	var keys []string
	for k := range gset.set {
		keys = append(keys, k)
	}
	return keys
}

// Returns true if element is included in set.
func (gset *GSet) Contains(key string) bool {
	_, ok := gset.set[key]
	return ok
}

type rawGSet []string

func (gset *GSet) MergeIn(other DData) error {
	o, ok := other.(*GSet)
	if !ok {
		return ErrInvalidTypeForMerge
	}
	for key, _ := range o.set {
		gset.set[key] = struct{}{}
	}
	return nil
}

func (gset *GSet) UnmarshalJSON(jsonBlob []byte) error {
	var raw rawGSet
	err := json.Unmarshal(jsonBlob, &raw)
	if err != nil {
		return err
	}
	for _, str := range raw {
		gset.set[str] = struct{}{}
	}
	return nil
}
func (gset *GSet) MarshalJSON() ([]byte, error) {
	res := gset.All()
	sort.Strings(res)
	return json.Marshal(res)
}
