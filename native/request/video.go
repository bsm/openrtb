package request

import (
	"encoding/json"

	"github.com/octoclick/openrtb-easyjson"
)

// Video is the native video object.
type Video struct {
	MIMEs       []string           `json:"mimes,omitempty"`       // Whitelist of content MIME types supported
	MinDuration int                `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration int                `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols   []openrtb.Protocol `json:"protocols,omitempty"`   // Video bid response protocols
	Ext         json.RawMessage    `json:"ext,omitempty"`
}
