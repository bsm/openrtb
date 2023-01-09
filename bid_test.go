package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestBid(t *testing.T) {
	var subject *Bid
	if err := fixture("bid", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Bid{
		ID:         "1",
		ImpID:      "1",
		Price:      0.751371,
		AdID:       "52a5516d29e435137c6f6e74",
		NoticeURL:  "http://ads.com/win/112770_1386565997?won=${AUCTION_PRICE}",
		AdMarkup:   "<html/>",
		AdvDomains: []string{"ads.com"},
		ImageURL:   "http://ads.com/112770_1386565997.jpeg",
		CampaignID: "52a5516d29e435137c6f6e74",
		CreativeID: "52a5516d29e435137c6f6e74_1386565997",
		DealID:     "example_deal",
		Attrs:      []CreativeAttribute{},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestBid_Validate(t *testing.T) {
	subject := &Bid{}
	if exp, got := ErrInvalidBidNoID, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Bid{ID: "BIDID"}
	if exp, got := ErrInvalidBidNoImpID, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
