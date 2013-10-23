package openrtb

// At least one of Bid is required.
// A bid response can contain multiple “seatbid” objects, each on behalf of a different bidder seat.
// Seatbid object can contain multiple bids each pertaining to a different impression on behalf of a seat.
// Each “bid” object must include the impression ID to which it pertains as well as the bid price.
// Group attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is
// only interested in winning any if it can win them all (i.e., all or nothing).
type Seatbid struct {
	Bid   []Bid      `json:"id"`             // Array of bid objects; each realtes to an imp, if exchange supported can have many bid objects.
	Seat  *string    `json:"seat,omiempty"`  // ID of the bidder seat optional string ID of the bidder seat on whose behalf this bid is made.
	Group *int       `json:"group,omiempty"` // '1' means impression must be won-lost as a group; default is '0'.
	Ext   Extensions `json:"ext,omiempty"`
}

// Validate Seatbid required attributes
func (sb *Seatbid) Valid() (bool, error) {

	if sb.Bid == nil || len(sb.Bid) < 1 {
		return false, errValidationSeatbidBid
	}

	for _, bid := range sb.Bid {
		if ok, err := bid.Valid(); !ok {
			return ok, err
		}
	}

	return true, nil
}
