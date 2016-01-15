package crdt

import (
	"errors"
)

type DData interface {
	MergeIn(other DData) error
	Unmarshal(interface{}) error
	Marshal() (interface{}, error)
}

var ErrInvalidTypeForMerge = errors.New("Invalid type for merge")
var ErrInvalidType = errors.New("Invalid type")
