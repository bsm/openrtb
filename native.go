package openrtb

// TODO: Add description, use cases and other details
type Native struct {
	Ver      *string    `json:"ver,omitempty"`      // Version of the Native Markup version in use
	Layout   *int       `json:"layout,omitempty"`   // Layout ID of the native ad unit
	Adunit   *int       `json:"adunit,omitempty"`   // Ad unit ID of the native ad unit
	Plcmtcnt *int       `json:"plcmtcnt,omitempty"` // Number of identical placements in this Layout
	Seq      *int       `json:"seq,omitempty"`      // This is not the sequence number of the content in the stream
	Assets   []Asset    `json:"assets"`             // Array of Asset Objects
	Ext      Extensions `json:"ext,omitempty"`
}

// Applies defaults
func (n *Native) WithDefaults() *Native {
	if n.Ver == nil {
		n.Ver = new(string)
		*n.Ver = "1"
	}
	if n.Plcmtcnt == nil {
		n.Plcmtcnt = new(int)
		*n.Plcmtcnt = 1
	}
	if n.Seq == nil {
		n.Seq = new(int)
		*n.Seq = 0
	}
	return n
}
