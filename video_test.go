package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Video", func() {
	var subject *Video

	BeforeEach(func() {
		Expect(fixture("video", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		boxingAllowed := 1
		Expect(subject).To(Equal(&Video{
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
		}))
	})

	It("should validate", func() {
		Expect((&Video{}).Validate()).To(Equal(ErrInvalidVideoNoMIMEs))
		Expect((&Video{
			MIMEs: []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoLinearity))
		Expect((&Video{
			MinDuration: 1,
			MaxDuration: 1,
			Linearity:   VideoLinearityNonLinear,
			MIMEs:       []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoProtocols))
		Expect((&Video{
			Protocol:    ProtocolVAST3,
			MinDuration: 1,
			MaxDuration: 1,
			Linearity:   VideoLinearityNonLinear,
			MIMEs:       []string{"video/mp4"},
		}).Validate()).NotTo(HaveOccurred())
	})
})
