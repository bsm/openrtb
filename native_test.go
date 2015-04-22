package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Native", func() {
	var subject *Native

	BeforeEach(func() {
		subject = new(Native)
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Ver).To(Equal("1"))
		Expect(*subject.Plcmtcnt).To(Equal(1))
		Expect(*subject.Seq).To(Equal(0))
	})
})
