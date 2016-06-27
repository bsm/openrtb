//go:generate ffjson $GOFILE
package openrtb

import "encoding/json"

// Private Marketplace Object
type Pmp struct {
	Private int              `json:"private_auction,omitempty"`
	Deals   []Deal           `json:"deals,omitempty"`
	Ext     *json.RawMessage `json:"ext,omitempty"`
}

// PMP Deal
type Deal struct {
	ID               string           `json:"id,omitempty"` // Unique deal ID
	BidFloor         float64          `json:"bidfloor,omitempty"`
	BidFloorCurrency string           `json:"bidfloorcur,omitempty"` // Currency of bid floor
	WSeat            []string         `json:"wseat,omitempty"`       // Array of buyer seats allowed to bid on this Direct Deal.
	WAdvDomain       []string         `json:"wadomain,omitempty"`    // Array of advertiser domains allowed to bid on this Direct Deal
	AuctionType      int              `json:"at,omitempty"`          // Optional override of the overall auction type of the bid request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in bidfloor is the agreed upon deal price. Additional auction types can be defined by the exchange.
	Ext              *json.RawMessage `json:"ext,omitempty"`

	Seats []string `json:"seats,omitempty"` // DEPRECATED: kept for backwards compatibility
	Type  int      `json:"type,omitempty"`  // DEPRECATED: kept for backwards compatibility
}

type jsonDeal Deal
