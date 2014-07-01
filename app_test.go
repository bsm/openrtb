package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var subject *App

	BeforeEach(func() {
		subject = new(App)
	})

	It("should have accessors", func() {
		Expect(subject.IsPrivacyPolicy()).To(BeFalse())
		Expect(subject.IsPaid()).To(BeFalse())
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Privacypolicy).To(Equal(0))
		Expect(*subject.Paid).To(Equal(0))
	})
})
