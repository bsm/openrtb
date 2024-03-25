package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb"
)

func TestMetric(t *testing.T) {
	var subject *Metric
	if err := fixture("metric", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Metric{
		Type:   "viewable_percentage",
		Value:  0.99,
		Vendor: "MOAT",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
