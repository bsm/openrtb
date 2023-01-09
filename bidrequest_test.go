package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestBidRequest(t *testing.T) {
	var subject *BidRequest
	if err := fixture("breq.banner", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	privacyPolicy := 1
	exp := &BidRequest{
		ID: "1234534625254",
		Impressions: []Impression{
			{
				ID:     "1",
				Secure: 1,
				Banner: &Banner{Width: 300, Height: 250, Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated}},
			},
		},
		Site: &Site{
			Inventory: Inventory{
				ID:            "234563",
				Name:          "Site ABCD",
				Domain:        "siteabcd.com",
				Categories:    []ContentCategory{ContentCategoryAutoParts, ContentCategoryAutoRepair},
				Publisher:     &Publisher{ID: "pub12345", Name: "Publisher A"},
				PrivacyPolicy: &privacyPolicy,
				Content: &Content{
					Keywords: "keyword a,keyword b,keyword c",
				},
			},
			Page:     "http://siteabcd.com/page.htm",
			Referrer: "http://referringsite.com/referringpage.htm",
		},
		Device: &Device{
			UA:           "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
			IP:           "64.124.253.1",
			OS:           "OS X",
			JS:           1,
			FlashVersion: "10.1",
		},
		User: &User{
			ID:       "45asdf987656789adfad4678rew656789",
			BuyerUID: "5df678asd8987656asdf78987654",
		},
		Test:              1,
		AuctionType:       2,
		TimeMax:           120,
		BlockedAdvDomains: []string{"company1.com", "company2.com"},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestBidRequest_complex(t *testing.T) {
	for _, kind := range []string{"exp", "video", "native"} {
		var subject *BidRequest
		if err := fixture("breq."+kind, &subject); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if err := subject.Validate(); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
}

func TestBidRequest_Validate(t *testing.T) {
	subject := &BidRequest{}
	if exp, got := ErrInvalidReqNoID, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &BidRequest{ID: "RID"}
	if exp, got := ErrInvalidReqNoImps, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &BidRequest{ID: "A", Impressions: []Impression{{ID: "1"}}, Site: &Site{}, App: &App{}}
	if exp, got := ErrInvalidReqMultiInv, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
