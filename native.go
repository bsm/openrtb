package openrtb

import (
	"bytes"
	"encoding/json"
)

// Native object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Native struct {
	Request      json.RawMessage     `json:"request"`         // Request payload complying with the Native Ad Specification.
	Version      string              `json:"ver,omitempty"`   // Version of the Native Ad Specification to which request complies; highly recommended for efficient parsing.
	APIs         []APIFramework      `json:"api,omitempty"`   // List of supported API frameworks for this impression.
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty"` // Blocked creative attributes
	Ext          json.RawMessage     `json:"ext,omitempty"`
}

type jsonNative Native

var (
	requestBytes = []byte("request")
	nullBytes    = []byte("null")
)

func (n *Native) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		return nil
	}

	if bytes.Contains(data, requestBytes) {
		var jn jsonNative
		if err := json.Unmarshal(data, &jn); err != nil {
			return err
		}

		*n = (Native)(jn)
		return nil
	}

	req := make(json.RawMessage, len(data))
	copy(req, data)
	n.Request = req
	return nil
}
