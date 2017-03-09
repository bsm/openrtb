package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner", func() {
	var subject *Banner

	BeforeEach(func() {
		err := fixture("banner", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Banner{
			W:     728,
			H:     90,
			Pos:   AdPosAboveFold,
			BType: []int{BannerTypeFrame},
			BAttr: []int{CreativeAttributeWindowsDialogOrAlert},
			Api:   []int{APIFrameworkMRAID1},
		}))
	})

})
