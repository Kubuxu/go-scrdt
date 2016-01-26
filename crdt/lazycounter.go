package crdt

import (
	"encoding/json"
)

// Lazy counter is not real CRDT. It is counter the will store
// biggest value observed.
// This means that conncurent increments do merge as maximum of one increment.
// Sometimes useful as it is basic and cheap type
type LCounter struct {
	value int
}

func NewLCounter() *LCounter {
	return &LCounter{0}
}

func (l *LCounter) Increment() {
	l.value++
}

func (l *LCounter) Value() int {
	return l.value
}

func (l *LCounter) MergeIn(other DData) error {
	o, ok := other.(*LCounter)
	if !ok {
		return ErrInvalidTypeForMerge
	}
	if o.value > l.value {
		l.value = o.value
	}
	return nil
}

func (l *LCounter) UnmarshalJSON(jsonBlob []byte) error {
	return json.Unmarshal(jsonBlob, &l.value)
}
func (l *LCounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.value)
}
