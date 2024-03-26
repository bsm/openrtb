package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb"
)

func TestSource(t *testing.T) {
	var subject *Source
	if err := fixture("source", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Source{
		FinalSaleDecision: 1,
		TransactionID:     "transaction-id",
		PaymentChain:      "payment-chain",
		Ext:               Extension("{}"),
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
