package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidRequest", func() {
	var subject *BidRequest
	privacyPolicy := 1

	BeforeEach(func() {
		Expect(fixture("breq.banner", &subject)).To(Succeed())
	})

	It("should parse complex requests", func() {
		for _, kind := range []string{"exp", "video", "native"} {
			var req *BidRequest
			err := fixture("breq."+kind, &req)
			Expect(err).NotTo(HaveOccurred(), "for %s", kind)
			Expect(req.Validate()).NotTo(HaveOccurred(), "for %s", kind)
		}
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&BidRequest{
			ID: "1234534625254",
			Impressions: []Impression{
				{
					ID:     "1",
					Secure: 1,
					Banner: &Banner{Width: 300, Height: 250, Position: AdPositionAboveFold, BlockedAttrs: []CreativeAttribute{CreativeAttributeUserInitiated}},
				},
			},
			Site: &Site{
				Inventory: Inventory{
					ID:            "234563",
					Name:          "Site ABCD",
					Domain:        "siteabcd.com",
					Categories:    []ContentCategory{ContentCategoryAutoParts, ContentCategoryAutoRepair},
					Publisher:     &Publisher{ID: "pub12345", Name: "Publisher A"},
					PrivacyPolicy: &privacyPolicy,
					Content: &Content{
						Keywords: "keyword a,keyword b,keyword c",
					},
				},
				Page:     "http://siteabcd.com/page.htm",
				Referrer: "http://referringsite.com/referringpage.htm",
			},
			Device: &Device{
				UA:           "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
				IP:           "64.124.253.1",
				OS:           "OS X",
				JS:           1,
				FlashVersion: "10.1",
			},
			User: &User{
				ID:       "45asdf987656789adfad4678rew656789",
				BuyerUID: "5df678asd8987656asdf78987654",
			},
			Test:              1,
			AuctionType:       2,
			TimeMax:           120,
			BlockedAdvDomains: []string{"company1.com", "company2.com"},
		}))
	})

	It("should validate", func() {
		Expect((&BidRequest{}).Validate()).To(Equal(ErrInvalidReqNoID))
		Expect((&BidRequest{ID: "A"}).Validate()).To(Equal(ErrInvalidReqNoImps))
		Expect((&BidRequest{ID: "A", Impressions: []Impression{{ID: "1"}}, Site: &Site{}, App: &App{}}).Validate()).To(Equal(ErrInvalidReqMultiInv))

		Expect((&BidRequest{ID: "A", Impressions: []Impression{{ID: "1", Banner: &Banner{}}}}).Validate()).NotTo(HaveOccurred())
		Expect((&BidRequest{ID: "A", Impressions: []Impression{{ID: "1", Banner: &Banner{}}}, Site: &Site{}}).Validate()).NotTo(HaveOccurred())
		Expect((&BidRequest{ID: "A", Impressions: []Impression{{ID: "1", Banner: &Banner{}}}, App: &App{}}).Validate()).NotTo(HaveOccurred())
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
