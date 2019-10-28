package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Audio", func() {
	var subject *Audio

	BeforeEach(func() {
		Expect(fixture("audio", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Audio{
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
		}))
	})

	It("should validate", func() {
		Expect((&Audio{
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
		}).Validate()).To(Equal(ErrInvalidAudioNoMIMEs))
	})

})
