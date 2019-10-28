package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Native", func() {
	var subject *Native

	BeforeEach(func() {
		Expect(fixture("native", &subject)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Native{
			Request: json.RawMessage(`"PAYLOAD"`),
			Version:     "2",
		}))
	})
})
