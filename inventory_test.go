package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var subject *App

	BeforeEach(func() {
		Expect(fixture("app", &subject)).To(Succeed())
	})

	It("should have accessors", func() {
		Expect(subject.GetPrivacyPolicy()).To(Equal(1))
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&App{
			Inventory: Inventory{
				ID:         "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfIUDA",
				Name:       "Yahoo Weather",
				Categories: []ContentCategory{"weather", ContentCategoryScience, ContentCategoryWeather},
				Publisher: &Publisher{
					ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
					Name:   "yahoo",
					Domain: "www.yahoo.com",
				},
			},
			Bundle:   "628677149",
			Version:  "1.0.2",
			StoreURL: "https://itunes.apple.com/id628677149",
		}))
	})
})

var _ = Describe("Site", func() {
	var subject *Site

	BeforeEach(func() {
		Expect(fixture("site", &subject)).To(Succeed())
	})

	It("should have accessors", func() {
		Expect(subject.GetPrivacyPolicy()).To(Equal(1))
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Site{
			Inventory: Inventory{
				ID:         "102855",
				Categories: []ContentCategory{ContentCategoryAdvertising},
				Domain:     "http://www.usabarfinder.com",
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
