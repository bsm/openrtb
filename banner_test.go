package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb/v3"
)

func TestBanner(t *testing.T) {
	var subject *Banner
	if err := fixture("banner", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Banner{
		W:     728,
		H:     90,
		Pos:   AdPosAboveFold,
		BType: []int{BannerTypeFrame},
		BAttr: []int{CreativeAttributeWindowsDialogOrAlert},
		Api:   []int{APIFrameworkMRAID1},
		VCM:   1,
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
