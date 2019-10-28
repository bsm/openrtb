package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidResponse", func() {
	var subject *BidResponse

	BeforeEach(func() {
		Expect(fixture("bres.single", &subject)).To(Succeed())
	})

	It("should parse complex responses", func() {
		for _, kind := range []string{"multi", "pmp", "vast"} {
			var req *BidResponse
			Expect(fixture("bres."+kind, &req)).To(Succeed(), "for %s", kind)
			Expect(req.Validate()).NotTo(HaveOccurred(), "for %s", kind)
		}
	})

	It("should parse responses", func() {
		Expect(subject).To(Equal(&BidResponse{
			ID: "BID-4-ZIMP-4b309eae-504a-4252-a8a8-4c8ceee9791a",
			SeatBids: []SeatBid{
				{
					Bids: []Bid{
						{
							ID:         "32a69c6ba388f110487f9d1e63f77b22d86e916b",
							ImpID:      "32a69c6ba388f110487f9d1e63f77b22d86e916b",
							Price:      0.065445,
							AdID:       "529833ce55314b19e8796116",
							NoticeURL:  "http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}",
							AdMarkup:   "<iframe src=\"foo.bar\"/>",
							AdvDomains: []string{},
							CampaignID: "529833ce55314b19e8796116",
							CreativeID: "529833ce55314b19e8796116_1385706446",
							Attrs:      []CreativeAttribute{},
						},
					},
					Seat: "772",
				},
			},
			Currency: "USD",
		}))
	})

	It("should validate", func() {
		Expect((&BidResponse{}).Validate()).To(Equal(ErrInvalidRespNoID))
		Expect((&BidResponse{ID: "RESPID"}).Validate()).To(Equal(ErrInvalidRespNoSeatBids))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})
})
