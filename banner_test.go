package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner", func() {
	var subject *Banner

	BeforeEach(func() {
		Expect(fixture("banner", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Banner{
			Width:        728,
			Height:       90,
			Position:     AdPositionAboveFold,
			BlockedTypes: []BannerType{BannerTypeFrame},
			BlockedAttrs: []CreativeAttribute{CreativeAttributeWindowsDialogOrAlert},
			APIs:         []APIFramework{APIFrameworkMRAID1},
			VCM:          1,
		}))
	})

})
