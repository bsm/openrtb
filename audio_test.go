package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestAudio(t *testing.T) {
	var subject *Audio
	if err := fixture("audio", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Audio{
		MIMEs: []string{
			"audio/mp4",
		},
		MinDuration:  5,
		MaxDuration:  30,
		Protocols:    []Protocol{ProtocolDAAST1, ProtocolDAAST1Wrapper},
		Sequence:     1,
		BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended:  30,
		MinBitrate:   300,
		MaxBitrate:   1500,
		Delivery:     []ContentDelivery{ContentDeliveryProgressive},
		CompanionAds: []Banner{
			{Width: 300, Height: 250, ID: "1234567893-1", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDirs: []ExpDir{ExpDirRight, ExpDirDown}},
			{Width: 728, Height: 90, ID: "1234567893-2", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		APIs:           []APIFramework{APIFrameworkVPAID1, APIFrameworkVPAID2},
		CompanionTypes: []CompanionType{CompanionTypeStatic, CompanionTypeHTML},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestAudio_Validate(t *testing.T) {
	subject := &Audio{
		MinDuration:  5,
		MaxDuration:  30,
		Protocols:    []Protocol{ProtocolDAAST1, ProtocolDAAST1Wrapper},
		Sequence:     1,
		BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended:  30,
		MinBitrate:   300,
		MaxBitrate:   1500,
		Delivery:     []ContentDelivery{ContentDeliveryProgressive},
		CompanionAds: []Banner{
			{Width: 300, Height: 250, ID: "1234567893-1", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDirs: []ExpDir{ExpDirRight, ExpDirDown}},
			{Width: 728, Height: 90, ID: "1234567893-2", Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		CompanionTypes: []CompanionType{CompanionTypeStatic, CompanionTypeHTML},
	}
	if exp, got := ErrInvalidAudioNoMIMEs, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
