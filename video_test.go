package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb/v3"
)

func TestVideo(t *testing.T) {
	var subject *Video
	if err := fixture("video", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	boxingAllowed := 1
	exp := &Video{
		Mimes: []string{
			"video/x-flv",
			"video/mp4",
			"application/x-shockwave-flash",
			"application/javascript",
		},
		MinDuration:    5,
		MaxDuration:    30,
		Protocols:      []int{VideoProtoVAST2, VideoProtoVAST3},
		W:              640,
		H:              480,
		Linearity:      VideoLinearityLinear,
		Sequence:       1,
		BAttr:          []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended:    30,
		MinBitrate:     300,
		MaxBitrate:     1500,
		BoxingAllowed:  &boxingAllowed,
		PlaybackMethod: []int{VideoPlaybackAutoSoundOn, VideoPlaybackClickToPlay},
		Delivery:       []int{ContentDeliveryProgressive},
		Pos:            AdPosAboveFold,
		CompanionAd: []Banner{
			{W: 300, H: 250, ID: "1234567893-1", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDir: []int{ExpDirRight, ExpDirDown}},
			{W: 728, H: 90, ID: "1234567893-2", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		Placement:     VideoPlacementInStream,
		Api:           []int{APIFrameworkVPAID1, APIFrameworkVPAID2},
		CompanionType: []int{VASTCompanionStatic, VASTCompanionHTML},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestVideo_Validate(t *testing.T) {
	subject := &Video{}
	if exp, got := ErrInvalidVideoNoMimes, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Video{Mimes: []string{"video/mp4"}}
	if exp, got := ErrInvalidVideoNoLinearity, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Video{
		MinDuration: 1,
		MaxDuration: 1,
		Linearity:   VideoLinearityNonLinear,
		Mimes:       []string{"video/mp4"},
	}
	if exp, got := ErrInvalidVideoNoProtocols, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
