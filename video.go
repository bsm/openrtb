package openrtb

import (
	"errors"
)

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string // Content MIME types supported.
	Linearity      *int     // Indicates whether the ad impression is linear or non-linear
	Minduration    *int     // Minimum video ad duration in seconds
	Maxduration    *int     // Maximum video ad duration in seconds
	Protocol       *int     // Video bid response protocols
	W              *int     // Width of the player in pixels
	H              *int     // Height of the player in pixels
	Startdelay     *int     // Indicates the start delay in seconds
	Sequence       *int     // Default: 1
	Battr          []int    // Blocked creative attributes
	Maxextended    *int     // Maximum extended video ad duration
	Minbitrate     *int     // Minimum bit rate in Kbps
	Maxbitrate     *int     // Maximum bit rate in Kbps
	Boxingallowed  *int     // If exchange publisher has rules preventing letter boxing
	Playbackmethod []int    // List of allowed playback methods
	Delivery       []int    // List of supported delivery methods
	Pos            *int     // Ad Position
	Companionad    []Banner
	Api            []int // List of supported API frameworks
	Companiontype  []int
	Ext            Extensions
}

// Returns the sequence number, with default fallback
func (v *Video) Seq() int {
	if v.Sequence != nil {
		return *v.Sequence
	}
	return 1
}

// Returns the boxing permission status, with default fallback
func (v *Video) IsBoxingAllowed() bool {
	if v.Boxingallowed != nil {
		return *v.Boxingallowed == 1
	}
	return true
}

// Returns the position, with default fallback
func (v *Video) Position() int {
	if v.Pos != nil {
		return *v.Pos
	}
	return AD_POS_UNKNOWN
}

// Validation errors
var (
	invalidVideoMimes       = errors.New("openrtb parse: video has no mimes")
	invalidVideoLinearity   = errors.New("openrtb parse: video linearity missing")
	invalidVideoMinduration = errors.New("openrtb parse: video minduration missing")
	invalidVideoMaxduration = errors.New("openrtb parse: video maxduration missing")
	invalidVideoProtocol    = errors.New("openrtb parse: video protocol missing")
)

// Validates the object
func (v *Video) Valid() (bool, error) {
	if len(v.Mimes) == 0 {
		return false, invalidVideoMimes
	} else if v.Linearity == nil {
		return false, invalidVideoLinearity
	} else if v.Minduration == nil {
		return false, invalidVideoMinduration
	} else if v.Maxduration == nil {
		return false, invalidVideoMaxduration
	} else if v.Protocol == nil {
		return false, invalidVideoProtocol
	}
	return true, nil
}

// Applies defaults
func (v *Video) WithDefaults() *Video {
	if v.Sequence == nil {
		v.Sequence = new(int)
		*v.Sequence = 1
	}
	if v.Boxingallowed == nil {
		v.Boxingallowed = new(int)
		*v.Boxingallowed = 1
	}
	if v.Pos == nil {
		v.Pos = new(int)
		*v.Pos = AD_POS_UNKNOWN
	}
	return v
}
