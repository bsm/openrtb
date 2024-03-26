package openrtb_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb"
)

func TestAudio(t *testing.T) {
	var subject *Audio
	if err := fixture("audio", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Audio{
		Mimes: []string{
			"audio/mp4",
		},
		MinDuration: 5,
		MaxDuration: 30,
		Protocols:   []int{AudioProtocolDAAST1, AudioProtocolDAAST1Wrapper},
		Sequence:    1,
		BAttr:       []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended: 30,
		MinBitrate:  300,
		MaxBitrate:  1500,
		Delivery:    []int{ContentDeliveryProgressive},
		CompanionAd: []Banner{
			{W: 300, H: 250, ID: "1234567893-1", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDir: []int{ExpDirRight, ExpDirDown}},
			{W: 728, H: 90, ID: "1234567893-2", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		API:           []int{APIFrameworkVPAID1, APIFrameworkVPAID2},
		CompanionType: []int{VASTCompanionStatic, VASTCompanionHTML},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestAudio_Validate(t *testing.T) {
	subject := &Audio{
		MinDuration: 5,
		MaxDuration: 30,
		Protocols:   []int{AudioProtocolDAAST1, AudioProtocolDAAST1Wrapper},
		Sequence:    1,
		BAttr:       []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert},
		MaxExtended: 30,
		MinBitrate:  300,
		MaxBitrate:  1500,
		Delivery:    []int{ContentDeliveryProgressive},
		CompanionAd: []Banner{
			{W: 300, H: 250, ID: "1234567893-1", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}, ExpDir: []int{ExpDirRight, ExpDirDown}},
			{W: 728, H: 90, ID: "1234567893-2", Pos: AdPosAboveFold, BAttr: []int{CreativeAttributeUserInitiated, CreativeAttributeWindowsDialogOrAlert}},
		},
		CompanionType: []int{VASTCompanionStatic, VASTCompanionHTML},
	}
	if exp, got := ErrInvalidAudioNoMimes, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
