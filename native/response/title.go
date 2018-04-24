package response

import "github.com/tisonet/openrtb-easyjson"

type Title struct {
	Text string            `json:"text"` // The text associated with the text element
	Ext  openrtb.Extension `json:"ext,omitempty"`
}
