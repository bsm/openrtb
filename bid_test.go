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

var _ = Describe("Loopme AR Bid", func() {
	var subject *Bid

	BeforeEach(func() {
		err := fixture("bid.loopme.ar", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Bid{
			ID:           "5d4b3b3be8250678c813f475",
			ImpID:        "1",
			Price:        29.7,
			AdID:         "2056515",
			NURL:         "https://us.tk0x1.com/sj/tr?et=BID_WIN&meta=MjA1NjUxNTo3Mjg0NDQ6NzFhY2I1ZDMtMmVkMi00MzdjLWJjYTItNjRiNDA4OWNkNjc0&ctx=CiQ3MWFjYjVkMy0yZWQyLTQzN2MtYmNhMi02NGI0MDg5Y2Q2NzQSAlJVGPy6LCDDwn0qCQgBEKDLmAEYCjjDwn1AAY8&ip=2.93.212.0&id=5d4b3b3be8250678c813f475&client_id=unity&a_ecp=29.7&a_price=${AUCTION_PRICE:BF}&name=imp_nurl",
			AdMarkup:     "<img>",
			AdvDomain:    []string{"unity.com"},
			IURL:         "https://loopmeedge.net/assets/2013906/23e11b6c-2690-e3c0-aa99-9efd38cd1fa6.jpg",
			CampaignID:   "2001626",
			CreativeID:   "2056515",
			Attr:         []int{99},
			Protocol:     2,
			Cat:          []string{"IAB19"},
			LURL:         "https://us.tk0x1.com/sj/tr?et=BID_LOSS&meta=MjA1NjUxNTo3Mjg0NDQ6NzFhY2I1ZDMtMmVkMi00MzdjLWJjYTItNjRiNDA4OWNkNjc0&ctx=CiQ3MWFjYjVkMy0yZWQyLTQzN2MtYmNhMi02NGI0MDg5Y2Q2NzQSAlJVGPy6LCDDwn0qCQgBEKDLmAEYCjjDwn1AAY8&ip=2.93.212.0&id=5d4b3b3be8250678c813f475&client_id=unity&a_ecp=29.7&a_id=${AUCTION_ID}&a_bid_id=${AUCTION_BID_ID}&a_loss_id=${AUCTION_LOSS}",
			H:            320,
			W:            480,
			CreativeType: "MRAID AR",
		}))
	})

	It("should validate", func() {
		Expect((&Bid{}).Validate()).To(Equal(ErrInvalidBidNoID))
		Expect((&Bid{ID: "BIDID"}).Validate()).To(Equal(ErrInvalidBidNoImpID))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
