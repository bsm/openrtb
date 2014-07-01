package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Seatbid", func() {
	var subject *Seatbid

	BeforeEach(func() {
		subject = new(Seatbid)
	})

	It("should have validation", func() {
		ok, err := subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		bid := (&Bid{}).SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)
		subject.Bid = append(subject.Bid, *bid)

		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

})
