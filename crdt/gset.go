package crdt

import (
	"sort"
)

type GSet struct {
	set map[string]struct{}
}

func NewGSet() *GSet {
	return &GSet{make(map[string]struct{})}
}

func (gset *GSet) Add(what string) {
	gset.set[what] = struct{}{}
}

func (gset *GSet) Len() int {
	return len(gset.set)
}

func (gset *GSet) All() []string {
	var keys []string
	for k := range gset.set {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (gset *GSet) Contains(key string) bool {
	_, ok := gset.set[key]
	return ok
}

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

func (gset *GSet) Unmarshal(input interface{}) error {
	in, ok := input.([]string)
	if !ok {
		return ErrInvalidType
	}
	for _, str := range in {
		gset.set[str] = struct{}{}
	}

	return nil
}
func (gset *GSet) Marshal() (interface{}, error) {
	return gset.All(), nil
}
