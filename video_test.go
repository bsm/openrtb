package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestVideo(t *testing.T) {
	var subject *Video
	if err := fixture("video", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	boxingAllowed := 1
	exp := &Video{
		MIMEs: []string{
			"video/x-flv",
			"video/mp4",
			"application/x-shockwave-flash",
			"application/javascript",
		},
		MinDuration:     5,
		MaxDuration:     30,
		Protocols:       []Protocol{ProtocolVAST2, ProtocolVAST3},
		Width:           640,
		Height:          480,
		Linearity:       VideoLinearityLinear,
		Sequence:        1,
		BlockedAttrs:    []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended:     30,
		MinBitrate:      300,
		MaxBitrate:      1500,
		BoxingAllowed:   &boxingAllowed,
		PlaybackMethods: []VideoPlayback{VideoPlaybackPageLoadSoundOn, VideoPlaybackClickToPlay},
		Delivery:        []ContentDelivery{ContentDeliveryProgressive},
		Position:        AdPositionAboveFold,
		CompanionAds: []Banner{
			{Width: 300, Height: 250, ID: "1234567893-1", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDirs: []ExpDir{ExpDirRight, ExpDirDown}},
			{Width: 728, Height: 90, ID: "1234567893-2", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		Placement:      VideoPlacementInStream,
		APIs:           []APIFramework{APIFrameworkVPAID1, APIFrameworkVPAID2},
		CompanionTypes: []CompanionType{CompanionTypeStatic, CompanionTypeHTML},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestVideo_Validate(t *testing.T) {
	subject := &Video{}
	if exp, got := ErrInvalidVideoNoMIMEs, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Video{MIMEs: []string{"video/mp4"}}
	if exp, got := ErrInvalidVideoNoLinearity, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
	subject = &Video{
		MinDuration: 1,
		MaxDuration: 1,
		Linearity:   VideoLinearityNonLinear,
		MIMEs:       []string{"video/mp4"},
	}
	if exp, got := ErrInvalidVideoNoProtocols, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
