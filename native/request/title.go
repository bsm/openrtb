package request

import "github.com/goccy/go-json"

// Title is the native title object.
type Title struct {
	Length int             `json:"len"` // Maximum length of the text in the title element
	Ext    json.RawMessage `json:"ext,omitempty"`
}
