package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Video", func() {
	var subject *Video

	BeforeEach(func() {
		err := fixture("video", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Video{
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
			BoxingAllowed:  iptr(1),
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
		}))
	})

	It("should validate", func() {
		Expect((&Video{}).Validate()).To(Equal(ErrInvalidVideoNoMimes))
		Expect((&Video{
			Mimes: []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoLinearity))
		Expect((&Video{
			Linearity: VideoLinearityNonLinear,
			Mimes:     []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoMinDuration))
		Expect((&Video{
			MinDuration: 1,
			Linearity:   VideoLinearityNonLinear,
			Mimes:       []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoMaxDuration))
		Expect((&Video{
			MinDuration: 1,
			MaxDuration: 1,
			Linearity:   VideoLinearityNonLinear,
			Mimes:       []string{"video/mp4"},
		}).Validate()).To(Equal(ErrInvalidVideoNoProtocols))
		Expect((&Video{
			Protocol:    VideoProtoVAST3,
			MinDuration: 1,
			MaxDuration: 1,
			Linearity:   VideoLinearityNonLinear,
			Mimes:       []string{"video/mp4"},
		}).Validate()).NotTo(HaveOccurred())
	})

})
