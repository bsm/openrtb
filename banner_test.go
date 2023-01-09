package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestBanner(t *testing.T) {
	var subject *Banner
	if err := fixture("banner", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Banner{
		Width:        728,
		Height:       90,
		Position:     AdPositionAboveFold,
		BlockedTypes: []BannerType{BannerTypeFrame},
		BlockedAttrs: []CreativeAttribute{CreativeAttributeWindowsDialogOrAlert},
		APIs:         []APIFramework{APIFrameworkMRAID1},
		VCM:          1,
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
