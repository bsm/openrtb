package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb"
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
		NURL:       "http://ads.com/win/112770_1386565997?won=${AUCTION_PRICE}",
		AdMarkup:   "<html/>",
		AdvDomain:  []string{"ads.com"},
		IURL:       "http://ads.com/112770_1386565997.jpeg",
		CampaignID: "52a5516d29e435137c6f6e74",
		CreativeID: "52a5516d29e435137c6f6e74_1386565997",
		DealID:     "example_deal",
		Attr:       []int{},
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

func TestLoopmeARBid(t *testing.T) {
	var subject *Bid
	if err := fixture("bid.loopme.ar", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Bid{
		ID:         "5d4b3b3be8250678c813f475",
		ImpID:      "1",
		Price:      29.7,
		AdID:       "2056515",
		NURL:       "https://us.fake-url.com/tr?et=BID_WIN&a_ecp=29.7&a_price=${AUCTION_PRICE:BF}&name=imp_nurl",
		AdMarkup:   "<img>",
		AdvDomain:  []string{"unity.com"},
		IURL:       "https://fake-url.net/assets/2690-e3c0-aa99-9efd38cd1fa6.jpg",
		CampaignID: "2001626",
		CreativeID: "2056515",
		Attr:       []int{99},
		Protocol:   2,
		Cat:        []string{"IAB19"},
		LURL:       "https://fake-url.com?et=BID_LOSS&meta=&a_id=${AUCTION_ID}&a_bid_id=${AUCTION_BID_ID}&a_loss_id=${AUCTION_LOSS}",
		H:          320,
		W:          480,
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
