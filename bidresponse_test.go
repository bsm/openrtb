package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestBidResponse(t *testing.T) {
	var subject *BidResponse
	if err := fixture("bres.single", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &BidResponse{
		ID: "BID-4-ZIMP-4b309eae-504a-4252-a8a8-4c8ceee9791a",
		SeatBids: []SeatBid{
			{
				Bids: []Bid{
					{
						ID:         "32a69c6ba388f110487f9d1e63f77b22d86e916b",
						ImpID:      "32a69c6ba388f110487f9d1e63f77b22d86e916b",
						Price:      0.065445,
						AdID:       "529833ce55314b19e8796116",
						NoticeURL:  "http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}",
						AdMarkup:   "<iframe src=\"foo.bar\"/>",
						AdvDomains: []string{},
						CampaignID: "529833ce55314b19e8796116",
						CreativeID: "529833ce55314b19e8796116_1385706446",
						Attrs:      []CreativeAttribute{},
					},
				},
				Seat: "772",
			},
		},
		Currency: "USD",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestBidResponse_complex(t *testing.T) {
	for _, kind := range []string{"multi", "pmp", "vast"} {
		var subject *BidResponse
		if err := fixture("bres."+kind, &subject); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if err := subject.Validate(); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
}

func TestBidResponse_Validate(t *testing.T) {
	subject := &BidResponse{}
	if exp, got := ErrInvalidRespNoID, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &BidResponse{ID: "RID"}
	if exp, got := ErrInvalidRespNoSeatBids, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
