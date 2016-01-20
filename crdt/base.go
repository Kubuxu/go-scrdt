//
// Copyright (c) 2016 Jakub Sztandera
// MIT Licensed; see the LICENSE file in this repository.
//

package crdt

import (
	"encoding/json"
	"errors"
)

type DData interface {
	json.Marshaler
	json.Unmarshaler
	// Merges 'other' into this object.
	MergeIn(other DData) error
}

var ErrInvalidTypeForMerge = errors.New("Invalid type for merge")
var ErrInvalidType = errors.New("Invalid type")
