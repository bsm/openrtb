package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidImpNoID        = errors.New("openrtb: impression ID missing")
	ErrInvalidImpMultiAssets = errors.New("openrtb: impression has multiple assets") // at least two out of Banner, Video, Native
)

// The "imp" object describes the ad position or impression being auctioned.  A single bid request
// can include multiple "imp" objects, a use case for which might be an exchange that supports
// selling all ad positions on a given page as a bundle.  Each "imp" object has a required ID so that
// bids can reference them individually.  An exchange can also conduct private auctions by
// restricting involvement to specific subsets of seats within bidders.
// The presence of Banner, Video, and/or Native objects
// subordinate to the Imp object indicates the type of impression being offered.
type Impression struct {
	ID                string         `json:"id"` // A unique identifier for this impression
	Metric            []*Metric      `json:"metric,omitempty"`
	Banner            *Banner        `json:"banner,omitempty"`
	Video             *Video         `json:"video,omitempty"`
	Audio             *Audio         `json:"audio,omitempty"`
	Native            *Native        `json:"native,omitempty"`
	Pmp               *Pmp           `json:"pmp,omitempty"`               // A reference to the PMP object containing any Deals eligible for the impression object.
	DisplayManager    string         `json:"displaymanager,omitempty"`    // Name of ad mediation partner, SDK technology, etc
	DisplayManagerVer string         `json:"displaymanagerver,omitempty"` // Version of the above
	Instl             int            `json:"instl,omitempty"`             // Interstitial, Default: 0 ("1": Interstitial, "0": Something else)
	TagID             string         `json:"tagid,omitempty"`             // IDentifier for specific ad placement or ad tag
	BidFloor          float64        `json:"bidfloor,omitempty"`          // Bid floor for this impression in CPM
	BidFloorCurrency  string         `json:"bidfloorcur,omitempty"`       // Currency of bid floor
	Secure            NumberOrString `json:"secure,omitempty"`            // Flag to indicate whether the impression requires secure HTTPS URL creative assets and markup.
	Exp               int            `json:"exp,omitempty"`               // Advisory as to the number of seconds that may elapse between the auction and the actual impression.
	IFrameBuster      []string       `json:"iframebuster,omitempty"`      // Array of names for supportediframe busters.
	ContentType       string         `json:"-"`                           // Used to attribute bid to a content - not sent in request
	Ext               Extension      `json:"ext,omitempty"`
}

func (imp *Impression) assetCount() int {
	n := 0
	if imp.Banner != nil {
		n++
	}
	if imp.Video != nil {
		n++
	}
	if imp.Native != nil {
		n++
	}
	return n
}

// Validates the `imp` object
func (imp *Impression) Validate() error {
	if imp.ID == "" {
		return ErrInvalidImpNoID
	}

	if count := imp.assetCount(); count > 1 {
		return ErrInvalidImpMultiAssets
	}

	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}

	return nil
}
