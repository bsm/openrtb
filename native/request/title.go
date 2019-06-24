package request

import "encoding/json"

type Title struct {
	Length int             `json:"len"` // Maximum length of the text in the title element
	Ext    json.RawMessage `json:"ext,omitempty"`
}
