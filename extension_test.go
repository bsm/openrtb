package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extension", func() {
	var _ json.Marshaler = (Extension)(nil)
	var _ json.Unmarshaler = (*Extension)(nil)

	It("should marshal JSON", func() {
		subject := Extension(`{"foo":"bar"}`)

		b, err := json.Marshal(subject)
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(Equal([]byte(`{"foo":"bar"}`)))
	})

	It("should unmarshal JSON", func() {
		subject := Extension([]byte{})

		err := json.Unmarshal([]byte(`{"foo":"bar"}`), &subject)
		Expect(err).NotTo(HaveOccurred())
		Expect(subject).To(Equal(Extension(`{"foo":"bar"}`)))
	})
})
