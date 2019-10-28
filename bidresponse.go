package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidRespNoID       = errors.New("openrtb: response missing ID")
	ErrInvalidRespNoSeatBids = errors.New("openrtb: response missing seatbids")
)

// BidResponse is the bid response wrapper object.
// ID and at least one "seatbid” object is required, which contains a bid on at least one impression.
// Other attributes are optional since an exchange may establish default values.
// No-Bids on all impressions should be indicated as a HTTP 204 response.
// For no-bids on specific impressions, the bidder should omit these from the bid response.
type BidResponse struct {
	ID         string          `json:"id"`                   // Reflection of the bid request ID for logging purposes
	SeatBids   []SeatBid       `json:"seatbid"`              // Array of seatbid objects
	BidID      string          `json:"bidid,omitempty"`      // Optional response tracking ID for bidders
	Currency   string          `json:"cur,omitempty"`        // Bid currency
	CustomData string          `json:"customdata,omitempty"` // Encoded user features
	NBR        NBR             `json:"nbr,omitempty"`        // Reason for not bidding, where 0 = unknown error, 1 = technical error, 2 = invalid request, 3 = known web spider, 4 = suspected Non-Human Traffic, 5 = cloud, data center, or proxy IP, 6 = unsupported device, 7 = blocked publisher or site, 8 = unmatched user
	Ext        json.RawMessage `json:"ext,omitempty"`        // Custom specifications in JSon
}

// Validate required attributes
func (res *BidResponse) Validate() error {
	if res.ID == "" {
		return ErrInvalidRespNoID
	} else if len(res.SeatBids) == 0 {
		return ErrInvalidRespNoSeatBids
	}

	for i := range res.SeatBids {
		sb := res.SeatBids[i]
		if err := (&sb).Validate(); err != nil {
			return err
		}
	}
	return nil
}
