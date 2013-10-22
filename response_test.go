package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_JSON(t *testing.T) {
	res := &Response{}
	res.Id = new(string)
	*res.Id = "RES_ID"
	res.Cur = new(string)
	*res.Cur = "USD"

	bytes, err := res.JSON()
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), `{"id":"RES_ID","seatbid":null,"cur":"USD"}`)
}

func TestResponse_Valid(t *testing.T) {
	res := &Response{}

	ok, err := res.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: missing ID")
	}

	res.Id = new(string)
	*res.Id = "RES_ID"

	ok, err = res.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb response: missing seatbids")
	}

	bid := &Bid{}
	bid.SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)

	sb := &Seatbid{}
	sb.Bid = append(sb.Bid, *bid)

	res.Seatbid = append(res.Seatbid, *sb)

	ok, err = res.Valid()
	assert.Equal(t, ok, true)
	assert.Nil(t, err)
}

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
