package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Source", func() {
	var subject *Source

	BeforeEach(func() {
		err := fixture("source", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Source{
			FinalSaleDecision: 1,
			TransactionID:     "transaction-id",
			PaymentChain:      "payment-chain",
			Ext:               Extension("{}"),
		}))
	})
})
