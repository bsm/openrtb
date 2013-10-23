package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeatbid_Valid(t *testing.T) {
	sb := &Seatbid{}

	ok, err := sb.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: seatbid is missing bids")
	}

	bid := &Bid{}
	bid.SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)
	sb.Bid = append(sb.Bid, *bid)

	ok, err = sb.Valid()
	assert.Equal(t, ok, true)
	assert.Nil(t, err)
}
