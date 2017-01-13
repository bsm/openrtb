package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Audio", func() {
	var subject *Audio

	BeforeEach(func() {
		err := fixture("audio", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Audio{
			Mimes: []string{
				"audio/mp4",
			},
			MinDuration: 5,
			MaxDuration: 30,
			Protocols:   []int{AudioProtocolDAAST1, AudioProtocolDAAST1Wrapper},
			Sequence:    1,
			BAttr:       []int{13, 14},
			MaxExtended: 30,
			MinBitrate:  300,
			MaxBitrate:  1500,
			Delivery:    []int{2},
			CompanionAd: []Banner{
				{W: 300, H: 250, ID: "1234567893-1", Pos: 1, BAttr: []int{13, 14}, ExpDir: []int{ExpDirRight, ExpDirDown}},
				{W: 728, H: 90, ID: "1234567893-2", Pos: 1, BAttr: []int{13, 14}},
			},
			CompanionType: []int{1, 2},
		}))
	})

	It("should validate", func() {
		Expect((&Audio{
			MinDuration: 5,
			MaxDuration: 30,
			Protocols:   []int{AudioProtocolDAAST1, AudioProtocolDAAST1Wrapper},
			Sequence:    1,
			BAttr:       []int{13, 14},
			MaxExtended: 30,
			MinBitrate:  300,
			MaxBitrate:  1500,
			Delivery:    []int{2},
			CompanionAd: []Banner{
				{W: 300, H: 250, ID: "1234567893-1", Pos: 1, BAttr: []int{13, 14}, ExpDir: []int{ExpDirRight, ExpDirDown}},
				{W: 728, H: 90, ID: "1234567893-2", Pos: 1, BAttr: []int{13, 14}},
			},
			CompanionType: []int{1, 2},
		}).Validate()).To(Equal(ErrInvalidAudioNoMimes))
	})

})
