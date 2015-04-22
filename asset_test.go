package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Asset", func() {
	var subject *Asset

	BeforeEach(func() {
		subject = new(Asset)
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Required).To(Equal(0))
	})
})
