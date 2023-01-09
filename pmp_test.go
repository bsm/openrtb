package openrtb_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestPMP(t *testing.T) {
	var subject *PMP
	if err := fixture("pmp", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &PMP{
		Private: 1,
		Deals: []Deal{
			{ID: "DX-1985-010A", BidFloor: 2.5, BidFloorCurrency: "", AuctionType: 2},
			{ID: "DX-1986-010A", BidFloor: 2.6, BidFloorCurrency: "", AuctionType: 2},
		},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestPMP_MarshalJSON(t *testing.T) {
	subject := &PMP{Deals: []Deal{{}}}
	exp := []byte(`{"deals":[{"at":2}]}`)
	if got, err := json.Marshal(subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	} else if !bytes.Equal(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
