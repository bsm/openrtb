package response

import "encoding/json"

// Title wraps title information.
type Title struct {
	Text string          `json:"text"` // The text associated with the text element
	Ext  json.RawMessage `json:"ext,omitempty"`
}
