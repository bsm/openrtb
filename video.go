package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidVideoNoMimes       = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity   = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoMinDuration = errors.New("openrtb: video min-duration missing")
	ErrInvalidVideoNoMaxDuration = errors.New("openrtb: video max-duration missing")
	ErrInvalidVideoNoProtocols   = errors.New("openrtb: video protocols missing")
)

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string  `json:"mimes,omitempty"`          // Content MIME types supported.
	MinDuration    *int      `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration    *int      `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Protocols      []int     `json:"protocols,omitempty"`      // Video bid response protocols
	Protocol       int       `json:"protocol,omitempty"`       // Video bid response protocols DEPRECATED
	W              int       `json:"w,omitempty"`              // Width of the player in pixels
	H              int       `json:"h,omitempty"`              // Height of the player in pixels
	StartDelay     int       `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Linearity      int       `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Skip           int       `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	SkipMin        int       `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter      int       `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence       int       `json:"sequence,omitempty"`       // Default: 1
	BAttr          []int     `json:"battr,omitempty"`          // Blocked creative attributes
	MaxExtended    int       `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	MinBitrate     int       `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	MaxBitrate     int       `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	BoxingAllowed  *int      `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	PlaybackMethod []int     `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery       []int     `json:"delivery,omitempty"`       // List of supported delivery methods
	Pos            int       `json:"pos,omitempty"`            // Ad Position
	CompanionAd    []Banner  `json:"companionad,omitempty"`
	Api            []int     `json:"api,omitempty"` // List of supported API frameworks
	CompanionType  []int     `json:"companiontype,omitempty"`
	Placement      int       `json:"placement,omitempty"` // Video placement type
	Ext            Extension `json:"ext,omitempty"`
}

type jsonVideo Video

// Validates the object
func (v *Video) Validate() error {
	if len(v.Mimes) == 0 {
		return ErrInvalidVideoNoMimes
	} else if v.Linearity == 0 {
		return ErrInvalidVideoNoLinearity
	} else if v.MinDuration == nil {
		return ErrInvalidVideoNoMinDuration
	} else if v.MaxDuration == nil {
		return ErrInvalidVideoNoMaxDuration
	} else if v.Protocol == 0 && len(v.Protocols) == 0 {
		return ErrInvalidVideoNoProtocols
	}
	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}

// MarshalJSON custom marshalling with normalization
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return json.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshalling with normalization
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Linearity == 0 {
		v.Linearity = VideoLinearityLinear
	}
}
