package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bid", func() {
	var subject *Bid

	BeforeEach(func() {
		err := fixture("bid", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Bid{
			ID:         "1",
			ImpID:      "1",
			Price:      0.751371,
			AdID:       "52a5516d29e435137c6f6e74",
			NURL:       "http://ads.com/win/112770_1386565997?won=${AUCTION_PRICE}",
			AdMarkup:   "<html/>",
			AdvDomain:  []string{"ads.com"},
			IURL:       "http://ads.com/112770_1386565997.jpeg",
			CampaignID: "52a5516d29e435137c6f6e74",
			CreativeID: "52a5516d29e435137c6f6e74_1386565997",
			DealID:     "example_deal",
			Attr:       []int{},
		}))
	})

	It("should validate", func() {
		Expect((&Bid{}).Validate()).To(Equal(ErrInvalidBidNoID))
		Expect((&Bid{ID: "BIDID"}).Validate()).To(Equal(ErrInvalidBidNoImpID))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
