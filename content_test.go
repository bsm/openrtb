package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Content", func() {
	var subject *Content

	BeforeEach(func() {
		err := fixture("content", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Content{
			ID:    "1",
			Title: "Video",
			Producer: &Producer{
				ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
				Name:   "yahoo",
				Domain: "www.yahoo.com",
			},
		}))
	})

})
