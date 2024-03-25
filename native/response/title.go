package response

import "github.com/UnityTech/openrtb"

// Title wraps title information.
type Title struct {
	Text string            `json:"text"` // The text associated with the text element
	Ext  openrtb.Extension `json:"ext,omitempty"`
}
