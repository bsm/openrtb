package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestQuantity(t *testing.T) {
	var subject *Quantity
	if err := fixture("quantity", &subject); err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}

	exp := &Quantity{
		Multiplier: 3.14,
		SourceType: MeasurementSourceTypePublisher,
	}

	if !reflect.DeepEqual(exp, subject) {
		t.Fatalf("expected %+v, got %+v", exp, subject)
	}
}

func TestQuantity_Validate(t *testing.T) {
	subject := Quantity{}
	if got := subject.Validate(); got != ErrMissingMultiplier {
		t.Fatalf("expected %+v, got %+v", ErrMissingMultiplier, got)
	}

	// If MeasurementSourceType is MeasurementSourceTypeMeasurementVendor, the
	// Vendor field should not be empty.
	subject = Quantity{
		Multiplier: 1.0,
		SourceType: MeasurementSourceTypeMeasurementVendor,
	}
	if got := subject.Validate(); got != ErrMissingMeasurementVendor {
		t.Fatalf("expected %+v, got %+v", ErrMissingMeasurementVendor, got)
	}

	subject.Vendor = "TestVendor" // Should fix the invalid Quantity
	if got := subject.Validate(); got != nil {
		t.Fatalf("expected no error, got %+v", got)
	}

	// All other value from MeasurementSourceType can be used without Vendor
	subject = Quantity{
		Multiplier: 2.0,
		SourceType: MeasurementSourceTypeUnknown,
	}
	if got := subject.Validate(); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
