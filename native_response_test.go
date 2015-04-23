package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NativeResponse", func() {
	var sptr = func(s string) *string { return &s }
	var iptr = func(i int) *int { return &i }

	var subject *NativeResponse

	BeforeEach(func() {
		subject = new(NativeResponse)
	})

	It("should have defaults", func() {
		subject.SetAssets(&ResponseAsset{})
		subject.WithDefaults()

		Expect(*subject.Ver).To(Equal("1"))
		Expect(*subject.Assets[0].Required).To(Equal(0))
	})

	It("should parse native adm", func() {
		resp, err := ParseNativeAdmBytes(testFixtures.simpleNativeResponse)

		Expect(err).NotTo(HaveOccurred())
		nativeAdm := NativeAdm{
			&NativeResponse{
				Ver: sptr("1"),
				Assets: []ResponseAsset{
					ResponseAsset{
						Id:       iptr(1),
						Required: iptr(0),
						Title:    &testFixtures.simpleTitle,
						Link:     &testFixtures.simpleLink,
					},
					ResponseAsset{
						Id:       iptr(2),
						Required: iptr(0),
						Data:     &testFixtures.simpleData,
					},
					ResponseAsset{
						Id:       iptr(3),
						Required: iptr(0),
						Img:      &testFixtures.simpleImg,
					},
					ResponseAsset{
						Id:       iptr(4),
						Required: iptr(0),
						Data:     &testFixtures.installData,
						Link:     &testFixtures.simpleLink,
					},
				},
				Link:        &testFixtures.fullLink,
				Imptrackers: []string{"http: //a.com/a", "http: //b.com/b"},
			},
		}

		Expect(resp).To(Equal(&nativeAdm))
	})
})
