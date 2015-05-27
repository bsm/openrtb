package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NativeRequest", func() {
	var subject *NativeRequest

	BeforeEach(func() {
		subject = new(NativeRequest)
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(*subject.Ver).To(Equal("1"))
		Expect(*subject.Plcmtcnt).To(Equal(1))
		Expect(*subject.Seq).To(Equal(0))
	})
})
