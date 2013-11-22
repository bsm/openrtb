package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_JSON(t *testing.T) {
	bid := &Bid{}
	bid.SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)

	sb := &Seatbid{}
	sb.Bid = append(sb.Bid, *bid)

	res := &Response{}
	res.Id = new(string)
	*res.Id = "RES_ID"
	res.Cur = new(string)
	*res.Cur = "USD"
	res.Seatbid = append(res.Seatbid, *sb)

	bytes, err := res.JSON()
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), `{"id":"RES_ID","seatbid":[{"bid":[{"id":"BIDID","impid":"IMPID","price":0}]}],"cur":"USD"}`)
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
