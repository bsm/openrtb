package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	var subject *Response

	BeforeEach(func() {
		subject = new(Response)
	})

	It("should validate", func() {
		ok, err := subject.Valid()
		Expect(err).To(Equal(ErrInvalidResID))
		Expect(ok).To(BeFalse())

		subject.Id = new(string)
		*subject.Id = "RES_ID"
		bid := (&Bid{}).SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)
		sb := Seatbid{}
		sb.Bid = append(sb.Bid, *bid)
		subject.Seatbid = append(subject.Seatbid, sb)

		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})
})
