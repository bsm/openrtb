package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Site", func() {
	var subject *Site

	BeforeEach(func() {
		subject = new(Site)
	})

	It("should have accessors", func() {
		Expect(subject.IsPrivacyPolicy()).To(BeFalse())
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Privacypolicy).To(Equal(0))
	})
})
