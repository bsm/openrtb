package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBid_Valid(t *testing.T) {
	bid := &Bid{}

	ok, err := bid.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: bid is missing ID")
	}

	bid.SetID("BIDID")
	ok, err = bid.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: bid is missing impression ID")
	}

	bid.SetImpID("IMPID")
	ok, err = bid.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: bid is missing price")
	}

	bid.SetPrice(0.0)
	ok, err = bid.Valid()
	assert.Equal(t, ok, true)
}
