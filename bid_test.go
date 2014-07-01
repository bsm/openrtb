package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bid", func() {
	var subject *Bid

	BeforeEach(func() {
		subject = new(Bid)
	})

	It("should have validation", func() {
		ok, err := subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		subject.SetID("BIDID")
		subject.SetImpID("IMPID")
		subject.SetPrice(0.0)

		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

})
