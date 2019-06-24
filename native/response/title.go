package response

import "encoding/json"

type Title struct {
	Text string          `json:"text"` // The text associated with the text element
	Ext  json.RawMessage `json:"ext,omitempty"`
}
