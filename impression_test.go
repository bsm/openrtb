package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impression", func() {
	var subject *Impression

	BeforeEach(func() {
		err := fixture("impression", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Impression{
			ID: "1",
			Metric: []*Metric{
				{
					Type:   "viewable_percentage_1",
					Value:  0.99,
					Vendor: "MOAT",
				},
				{
					Type:   "viewable_percentage_2",
					Value:  0.98,
					Vendor: "MOAT",
				},
			},
			Banner: &Banner{
				W: 300,
				H: 250,
			},
			BidFloor: 0.03,
			Pmp: &Pmp{
				Private: 1,
				Deals: []Deal{
					{
						ID:          "DX-1985-010A",
						BidFloor:    2.5,
						AuctionType: 2,
					},
				},
			},
			Rewarded: 0,
		}))
	})

	It("should validate", func() {
		Expect((&Impression{}).Validate()).To(Equal(ErrInvalidImpNoID))
		Expect((&Impression{ID: "IMPID", Banner: &Banner{}, Video: &Video{}}).Validate()).To(Equal(ErrInvalidImpMultiAssets))
		Expect((&Impression{ID: "IMPID", Banner: &Banner{}}).Validate()).NotTo(HaveOccurred())
	})

})
