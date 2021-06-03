package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bid", func() {
	var subject *Bid

	BeforeEach(func() {
		Expect(fixture("bid", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Bid{
			ID:         "1",
			ImpID:      "1",
			Price:      0.751371,
			AdID:       "52a5516d29e435137c6f6e74",
			NoticeURL:  "http://ads.com/win/112770_1386565997?won=${AUCTION_PRICE}",
			AdMarkup:   "<html/>",
			AdvDomains: []string{"ads.com"},
			ImageURL:   "http://ads.com/112770_1386565997.jpeg",
			CampaignID: "52a5516d29e435137c6f6e74",
			CreativeID: "52a5516d29e435137c6f6e74_1386565997",
			DealID:     "example_deal",
			Attrs:      []CreativeAttribute{},
		}))
	})

	It("should parse asian correctly", func() {
		var asian *Bid
		Expect(fixture("asian_bid", &asian)).To(Succeed())

		expectNative := `{"link":{"url":"https://test.com"},"assets":[{"id":1,"required":1,"title":{"text":"TEST!"}}]}`
		Expect(asian).To(Equal(&Bid{
			ID:     "BID-SCI-258-471500-k5p21f06-1dm3",
			ImpID:  "SCI-1-471500-k5p21f06-1dm3",
			Price:  0.0001,
			Native: []byte(expectNative),
		}))
	})

	It("should validate", func() {
		Expect((&Bid{}).Validate()).To(Equal(ErrInvalidBidNoID))
		Expect((&Bid{ID: "BIDID"}).Validate()).To(Equal(ErrInvalidBidNoImpID))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
