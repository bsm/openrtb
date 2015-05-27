package openrtb

// The presence of a Native as a subordinate of the Imp object indicates that this impression
// is offered as a native type impression.
// The Native Subcommittee has developed a companion specification to OpenRTB called the Native
// Ad Specification. It defines the request parameters and response markup structure of native
// ad units.
// This object provides the means of transporting request parameters as an opaque string
type Native struct {
	Request *string    `json:"request"`         // Request payload complying with the Native Ad Specification
	Ver     *string    `json:"ver,omitempty"`   // Version of the Native Ad Specification to which request complies
	Api     []int      `json:"api,omitempty"`   // List of supported API frameworks for this impression
	Battr   []int      `json:"battr,omitempty"` // Blocked creative attributes
	Ext     Extensions `json:"ext,omitempty"`
}
