package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impression", func() {
	var subject *Impression

	BeforeEach(func() {
		subject = new(Impression)
	})

	It("should validate", func() {
		b := &Banner{}
		v := (&Video{Mimes: []string{"MIME_123"}}).SetLinearity(1).SetMinduration(1).SetMaxduration(5).SetProtocol(1)

		ok, err := subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		subject.SetId("CODE_12")
		subject.SetBanner(b)
		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())

		subject.SetVideo(v)
		ok, err = subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		subject.SetBanner(nil)
		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

	It("should have accessors", func() {
		Expect(subject.IsSecure()).To(BeFalse())
	})

	It("should have defaults", func() {
		subject.SetBanner(&Banner{}).SetVideo(&Video{})
		subject.WithDefaults()

		Expect(*subject.Instl).To(Equal(0))
		Expect(*subject.Secure).To(Equal(0))
		Expect(*subject.Bidfloor).To(Equal(float32(0.0)))
		Expect(*subject.Bidfloorcur).To(Equal("USD"))
		Expect(*subject.Banner.Topframe).To(Equal(0))
		Expect(*subject.Banner.Pos).To(Equal(AD_POS_UNKNOWN))
		Expect(*subject.Video.Sequence).To(Equal(1))
		Expect(*subject.Video.Boxingallowed).To(Equal(1))
		Expect(*subject.Video.Pos).To(Equal(AD_POS_UNKNOWN))
	})
})
