package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var subject *App

	BeforeEach(func() {
		err := fixture("app", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should have accessors", func() {
		Expect(subject.GetPrivacyPolicy()).To(Equal(1))
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&App{
			Inventory: Inventory{
				ID:   "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfIUDA",
				Name: "Yahoo Weather",
				Cat:  []string{"weather", "IAB15", "IAB15-10"},
				Publisher: &Publisher{
					ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
					Name:   "yahoo",
					Domain: "www.yahoo.com",
				},
			},
			Bundle:   "628677149",
			Ver:      "1.0.2",
			StoreURL: "https://itunes.apple.com/id628677149",
		}))
	})

})

var _ = Describe("Site", func() {
	var subject *Site

	BeforeEach(func() {
		err := fixture("site", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should have accessors", func() {
		Expect(subject.GetPrivacyPolicy()).To(Equal(1))
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Site{
			Inventory: Inventory{
				ID:     "102855",
				Cat:    []string{"IAB3-1"},
				Domain: "http://www.usabarfinder.com",
				Publisher: &Publisher{
					ID:     "8953",
					Name:   "local.com",
					Domain: "local.com",
				},
			},
			Page: "http://eas.usabarfinder.com/eas?cu=13824;cre=mu;target=_blank",
		}))
	})

})
