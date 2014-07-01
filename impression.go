package openrtb

import (
	"errors"
)

// The "imp" object describes the ad position or impression being auctioned.  A single bid request
// can include multiple "imp" objects, a use case for which might be an exchange that supports
// selling all ad positions on a given page as a bundle.  Each "imp" object has a required ID so that
// bids can reference them individually.  An exchange can also conduct private auctions by
// restricting involvement to specific subsets of seats within bidders.
type Impression struct {
	Id                *string    `json:"id` // A unique identifier for this impression
	Banner            *Banner    `json:"banner,omitempty"`
	Video             *Video     `json:"video,omitempty"`
	Displaymanager    *string    `json:"displaymanager,omitempty"`    // Name of ad mediation partner, SDK technology, etc
	Displaymanagerver *string    `json:"displaymanagerver,omitempty"` // Version of the above
	Instl             *int       `json:"instl,omitempty"`             // Interstitial, Default: 0 ("1": Interstitial, "0": Something else)
	Tagid             *string    `json:"tagid,omitempty"`             // Identifier for specific ad placement or ad tag
	Bidfloor          *float32   `json:"bidfloor,omitempty"`          // Bid floor for this impression in CPM
	Bidfloorcur       *string    `json:"bidfloorcur,omitempty"`       // Currency of bid floor
	Secure            *int       `json:"secure,omitempty"`            // Flag to indicate whether the impression requires secure HTTPS URL creative assets and markup.
	Iframebuster      []string   `json:"iframebuster,omitempty"`      // Array of names for supportediframe busters.
	Pmp               *Pmp       `json:"pmp,omitempty"`               // A reference to the PMP object containing any Deals eligible for the impression object.
	Ext               Extensions `json:"ext,omitempty"`
}

// Validation errors
var (
	ErrInvalidImpID  = errors.New("openrtb parse: impression ID missing")
	ErrInvalidImpBoV = errors.New("openrtb parse: impression has neither a banner nor video")
	ErrInvalidImpBaV = errors.New("openrtb parse: impression has banner and video")
)

// Validates the `imp` object
func (imp *Impression) Valid() (bool, error) {

	if imp.Id == nil {
		return false, ErrInvalidImpID
	} else if imp.Banner != nil && imp.Video != nil {
		return false, ErrInvalidImpBaV
	} else if imp.Video != nil {
		if ok, err := imp.Video.Valid(); !ok {
			return ok, err
		}
	} else if imp.Banner == nil {
		return false, ErrInvalidImpBoV
	}

	return true, nil
}

// Returns secure status, with default fallback
func (imp *Impression) IsSecure() bool {
	if imp.Secure != nil {
		return *imp.Secure == 1
	}
	return false
}

// Returns the `imp` object returning defaults
func (imp *Impression) WithDefaults() *Impression {
	if imp.Instl == nil {
		imp.Instl = new(int)
		*imp.Instl = 0
	}

	if imp.Bidfloor == nil {
		imp.Bidfloor = new(float32)
		*imp.Bidfloor = 0
	}

	if imp.Secure == nil {
		imp.Secure = new(int)
		*imp.Secure = 0
	}

	if imp.Bidfloorcur == nil {
		imp.Bidfloorcur = new(string)
		*imp.Bidfloorcur = "USD"
	}

	if imp.Banner != nil {
		imp.Banner.WithDefaults()
	}

	if imp.Video != nil {
		imp.Video.WithDefaults()
	}

	return imp
}

// Set the ID
func (imp *Impression) SetId(id string) *Impression {
	if imp.Id == nil {
		imp.Id = new(string)
	}
	*imp.Id = id
	return imp
}

// Set the Banner
func (imp *Impression) SetBanner(b *Banner) *Impression {
	imp.Banner = b
	return imp
}

// Set the Video
func (imp *Impression) SetVideo(v *Video) *Impression {
	imp.Video = v
	return imp
}
