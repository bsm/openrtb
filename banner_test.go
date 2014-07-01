package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner", func() {
	var subject *Banner

	BeforeEach(func() {
		subject = new(Banner)
	})

	It("should have accessors", func() {
		Expect(subject.IsTopFrame()).To(BeFalse())
		Expect(subject.Position()).To(Equal(0))
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Topframe).To(Equal(0))
		Expect(*subject.Pos).To(Equal(AD_POS_UNKNOWN))
	})
})
