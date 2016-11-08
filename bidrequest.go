//go:generate ffjson $GOFILE
package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidReqNoID     = errors.New("openrtb: request ID missing")
	ErrInvalidReqNoImps   = errors.New("openrtb: request has no impressions")
	ErrInvalidReqMultiInv = errors.New("openrtb: request has multiple inventory sources") // has site and app
)

// The top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type BidRequest struct {
	ID          string          `json:"id"` // Unique ID of the bid request
	Imp         []Impression    `json:"imp,omitempty"`
	Site        Site            `json:"site,omitempty"`
	App         App             `json:"app,omitempty"`
	Device      Device          `json:"device,omitempty"`
	User        User            `json:"user,omitempty"`
	Test        int             `json:"test,omitempty"`    // Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode
	AuctionType int             `json:"at"`                // Auction type, where 1 = First Price, 2 = Second Price Plus. Exchange-specific auction types can be defined using values greater than 500.
	TMax        int             `json:"tmax,omitempty"`    // Maximum amount of time in milliseconds to submit a bid
	WSeat       []string        `json:"wseat,omitempty"`   // Array of buyer seats allowed to bid on this auction
	AllImps     int             `json:"allimps,omitempty"` // Flag to indicate whether exchange can verify that all impressions offered represent all of the impressions available in context, Default: 0
	Cur         []string        `json:"cur,omitempty"`     // Array of allowed currencies
	Bcat        []string        `json:"bcat,omitempty"`    // Blocked Advertiser Categories.
	BAdv        []string        `json:"badv,omitempty"`    // Array of strings of blocked toplevel domains of advertisers
	Regs        Regulations     `json:"regs,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`

	Pmp *Pmp `json:"pmp,omitempty"` // DEPRECATED: kept for backwards compatibility
}

// Validates the request
func (req *BidRequest) Validate() error {
	if req.ID == "" {
		return ErrInvalidReqNoID
	} else if len(req.Imp) == 0 {
		return ErrInvalidReqNoImps
	} else if req.Site.ID != "" && req.App.ID != "" {
		return ErrInvalidReqMultiInv
	}

	for _, imp := range req.Imp {
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}

	return nil
}
