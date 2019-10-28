package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PMP", func() {
	var subject *PMP

	BeforeEach(func() {
		Expect(fixture("pmp", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&PMP{
			Private: 1,
			Deals: []Deal{
				{ID: "DX-1985-010A", BidFloor: 2.5, BidFloorCurrency: "", AuctionType: 2},
				{ID: "DX-1986-010A", BidFloor: 2.6, BidFloorCurrency: "", AuctionType: 2},
			},
		}))
	})

	It("should generate correctly", func() {
		Expect(json.Marshal(&PMP{Deals: []Deal{{}}})).To(MatchJSON(`{"deals":[{"at":2}]}`))
	})

})
