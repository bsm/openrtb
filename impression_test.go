package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestImpression(t *testing.T) {
	var subject *Impression
	if err := fixture("impression", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Impression{
		ID: "1",
		Banner: &Banner{
			Width:  300,
			Height: 250,
		},
		BidFloor: 0.03,
		PMP: &PMP{
			Private: 1,
			Deals: []Deal{
				{
					ID:          "DX-1985-010A",
					BidFloor:    2.5,
					AuctionType: 2,
				},
			},
		},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestImpression_Validate(t *testing.T) {
	subject := &Impression{}
	if exp, got := ErrInvalidImpNoID, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Impression{ID: "IMPID", Banner: &Banner{}, Video: &Video{}}
	if exp, got := ErrInvalidImpMultiAssets, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
