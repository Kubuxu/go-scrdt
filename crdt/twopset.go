package crdt

type TwoPSet struct {
	add *GSet
	remove *GSet
}

func NewTwoPSet() *TwoPSet {
	return &TwoPSet{NewGSet(), NewGSet()}
}

func (set *TwoPSet) Add(what string) {
	set.add.Add(what)
}

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

func (set *TwoPSet) Contains(key string) bool {
	return set.add.Contains(key) && (!set.remove.Contains(key))
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
type twoPMarshal struct {
	A []string
	R []string
}

func (set *TwoPSet) Unmarshal(input interface{}) error {
	in, ok := input.(twoPMarshal)
	if !ok {
		return ErrInvalidType
	}
	set.add.Unmarshal(in.A)
	set.remove.Unmarshal(in.R)

	return nil
}
func (set *TwoPSet) Marshal() (interface{}, error) {
	A, _ := set.add.Marshal()
	R, _ := set.remove.Marshal()
	return twoPMarshal{A.([]string), R.([]string)}, nil
}

