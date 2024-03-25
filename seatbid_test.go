package openrtb_test

import (
	"errors"
	"testing"

	. "github.com/UnityTech/openrtb"
)

func TestSeatBid_Validate(t *testing.T) {
	subject := &SeatBid{}
	if exp, got := ErrInvalidSeatBidBid, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
