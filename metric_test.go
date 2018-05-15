package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metric", func() {
	var metric *Metric

	BeforeEach(func() {
		err := fixture("metric", &metric)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(metric).To(Equal(&Metric{
			Type:   "viewable_percentage",
			Value:  0.99,
			Vendor: "MOAT",
		}))
	})

})
