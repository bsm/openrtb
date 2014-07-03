package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Video", func() {
	var subject *Video

  BeforeEach(func() { subject = new(Video) })

  Describe("Seq()", func() {
    It("should return '1' as default", func() {
      Expect(subject.Seq()).To(Equal(1))
    })
  })

  Describe("IsBoxingAllowed()", func() {
    It("should return true as default", func() {
      Expect(subject.IsBoxingAllowed()).To(BeTrue())
    })

    It("should return false when set as false", func() {
      subject.Boxingallowed = new(int)
      *subject.Boxingallowed = 0
      Expect(subject.IsBoxingAllowed()).To(BeFalse())
    })
  })

  Describe("Position()", func() {
    It("should return 0/AD_POST_UNKNOWN as default", func() {
      Expect(subject.Position()).To(Equal(0))
    })
  })

  Describe("Valid()", func() {
    It("should return error messages when attributes missing", func() {
    	ok, err := subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video has no mimes"))

      subject.Mimes  = []string{"RAND_KEY"} // With Mimes
    	ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video linearity missing"))

	  	subject.SetLinearity(2)   // With Linearity
    	ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video minduration missing"))

	  	subject.SetMinduration(1) // With Minduration
    	ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video maxduration missing"))

	  	subject.SetMaxduration(5) // With Maxduration
    	ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video protocol missing"))

	  	subject.SetProtocol(1)   // With Protocol
      ok, err = subject.Valid()
      Expect(err).NotTo(HaveOccurred())
    	Expect(ok).To(BeTrue())
    })
  })

  Describe("WithDefaults()", func() {

    It("should return object with default values", func() {
      subject.WithDefaults()
      Expect(*subject.Sequence).To(Equal(1))
      Expect(*subject.Boxingallowed).To(Equal(1))
      Expect(*subject.Pos).To(Equal(AD_POS_UNKNOWN))
    })
  })

})