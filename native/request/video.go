package request

import "github.com/Upliner/openrtb"

// TODO unclear if its the same as imp.video https://github.com/openrtb/OpenRTB/issues/26
type Video struct {
	Mimes       []string          `json:"mimes,omitempty"`       // Whitelist of content MIME types supported
	MinDuration int               `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration int               `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols   []int             `json:"protocols,omitempty"`   // Video bid response protocols
	Ext         openrtb.Extension `json:"ext,omitempty"`
}
