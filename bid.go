//go:generate ffjson $GOFILE
package openrtb

import "errors"

// Validation errors
var (
	ErrInvalidBidNoID    = errors.New("openrtb: bid is missing ID")
	ErrInvalidBidNoImpID = errors.New("openrtb: bid is missing impression ID")
)

// ID, ImpID and Price are required; all other optional.
// If the bidder wins the impression, the exchange calls notice URL (nurl)
// a) to inform the bidder of the win;
// b) to convey certain information using substitution macros.
// Adomain can be used to check advertiser block list compliance.
// Cid can be used to block ads that were previously identified as inappropriate.
// Substitution macros may allow a bidder to use a static notice URL for all of its bids.
type Bid struct {
	ID         string   `json:"id"`
	ImpID      string   `json:"impid"`             // Required string ID of the impression object to which this bid applies.
	Price      float64  `json:"price"`             // Bid price in CPM. Suggests using integer math for accounting to avoid rounding errors.
	AdID       string   `json:"adid,omitempty"`    // References the ad to be served if the bid wins.
	NURL       string   `json:"nurl,omitempty"`    // Win notice URL.
	AdMarkup   string   `json:"adm,omitempty"`     // Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	AdvDomain  []string `json:"adomain,omitempty"` // Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	IURL       string   `json:"iurl,omitempty"`    // Sample image URL.
	CampaignID string   `json:"cid,omitempty"`     // Campaign ID that appears with the Ad markup.
	CreativeID string   `json:"crid,omitempty"`    // Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Cat        []string `json:"cat,omitempty"`     // IAB content categories of the creative. Refer to List 5.1
	Attr       []int    `json:"attr,omitempty"`    // Array of creative attributes.
	DealID     string   `json:"dealid,omitempty"`  // DealID extension of private marketplace deals
	H          int      `json:"h,omitempty"`       // Height of the ad in pixels.
	W          int      `json:"w,omitempty"`       // Width of the ad in pixels.
	Ext        BidExt   `json:"ext,omitempty"`
}

type BidExt struct {
	ImpressionTrackingUrl []string `json:"impression_tracking_url,omitempty"`
	Duration              int64    `json:"duration"`
	VastUrl               string   `json:"vast_url"`
}

// Validate required attributes
func (bid *Bid) Validate() error {
	if bid.ID == "" {
		return ErrInvalidBidNoID
	} else if bid.ImpID == "" {
		return ErrInvalidBidNoImpID
	}

	return nil
}
