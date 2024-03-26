package openrtb

import (
	"encoding/json"
	"errors"
)

// Extension is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler, defined in encoding/json package
type Extension json.RawMessage

// MarshalJSON returns e as the JSON encoding of e.
func (e Extension) MarshalJSON() ([]byte, error) {
	return json.RawMessage(e).MarshalJSON()
}

// UnmarshalJSON sets *e to a copy of data.
func (e *Extension) UnmarshalJSON(data []byte) error {
	if e == nil {
		return errors.New("openrtb.Extension: UnmarshalJSON on nil pointer")
	}
	return (*json.RawMessage)(e).UnmarshalJSON(data)
}
