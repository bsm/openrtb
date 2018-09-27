package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Content", func() {
	var subject *Content

	BeforeEach(func() {
		err := fixture("content", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Content{
			ID:    "1",
			Title: "Video",
			Producer: &Producer{
				ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
				Name:   "yahoo",
				Domain: "www.yahoo.com",
			},
		}))
	})

})

var _ = Describe("QuotedContent", func() {
	var subject *Content

	BeforeEach(func() {
		err := fixture("content.quoted", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Content{
			Keywords:           "Orwell, 1984",
			UserRating:         "3",
			Episode:            1,
			Title:              "This is the video title",
			URL:                "http://cdnp.tremormedia.com/video/1984.flv",
			LiveStream:         0,
			ContentRating:      "G",
			Len:                129,
			Series:             "book reading",
			VideoQuality:       2,
			Context:            1,
			Season:             "1",
			SourceRelationship: 0,
			ID:                 "eb9f13ede5fd225333971523f6042f9d",
		}))
	})

})
