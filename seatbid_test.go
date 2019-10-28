package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SeatBid", func() {
	It("should validate", func() {
		Expect((&SeatBid{}).Validate()).To(Equal(ErrInvalidSeatBidBid))
		Expect((&SeatBid{Bids: []Bid{
			{ID: "BIDID", ImpID: "IMPID"}},
		}).Validate()).NotTo(HaveOccurred())
	})
})
