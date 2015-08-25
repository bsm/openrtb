package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pmp", func() {
	var subject *Pmp

	BeforeEach(func() {
		err := fixture("pmp", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Pmp{
			Private: 1,
			Deals: []Deal{
				{ID: "DX-1985-010A", BidFloor: 2.5, BidFloorCurrency: "", AuctionType: 2},
				{ID: "DX-1986-010A", BidFloor: 2.6, BidFloorCurrency: "", AuctionType: 2},
			},
		}))
	})

	It("should generate correctly", func() {
		bin, err := json.Marshal(&Pmp{Deals: []Deal{{}}})
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bin)).To(Equal(`{"deals":[{"at":2}]}`))
	})

})
