package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberOrString", func() {

	It("should decode numbers", func() {
		var n NumberOrString
		Expect(json.Unmarshal([]byte(`33`), &n)).To(Succeed())
		Expect(n).To(Equal(NumberOrString(33)))
	})

	It("should decode strings", func() {
		var n NumberOrString
		Expect(json.Unmarshal([]byte(`"33"`), &n)).To(Succeed())
		Expect(n).To(Equal(NumberOrString(33)))
	})

	It("should encode to numbers", func() {
		var n NumberOrString = 33
		bin, err := json.Marshal(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bin)).To(Equal(`33`))
	})

})

var _ = Describe("StringOrNumber", func() {

	It("should decode numbers", func() {
		var n StringOrNumber
		Expect(json.Unmarshal([]byte(`33`), &n)).To(Succeed())
		Expect(n).To(Equal(StringOrNumber("33")))
	})

	It("should decode strings", func() {
		var n StringOrNumber
		Expect(json.Unmarshal([]byte(`"33"`), &n)).To(Succeed())
		Expect(n).To(Equal(StringOrNumber("33")))
	})

	It("should encode to strings", func() {
		var n StringOrNumber = "33"
		bin, err := json.Marshal(n)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bin)).To(Equal(`"33"`))
	})

})
