package openrtb

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	var subject *Response
	var sptr = func(s string) *string { return &s }
	var fptr = func(f float32) *float32 { return &f }

	BeforeEach(func() {
		subject = new(Response)
	})

	It("should validate", func() {
		ok, err := subject.Valid()
		Expect(err).To(Equal(ErrInvalidResID))
		Expect(ok).To(BeFalse())

		subject.Id = new(string)
		*subject.Id = "RES_ID"
		bid := (&Bid{}).SetID("BIDID").SetImpID("IMPID").SetPrice(0.0)
		sb := Seatbid{}
		sb.Bid = append(sb.Bid, *bid)
		subject.Seatbid = append(subject.Seatbid, sb)

		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

	It("should parse responses", func() {
		resp, err := ParseResponse(bytes.NewBuffer(testFixtures.simpleResponse))
		Expect(err).NotTo(HaveOccurred())

		bid := Bid{
			Id:      sptr("32a69c6ba388f110487f9d1e63f77b22d86e916b"),
			Impid:   sptr("32a69c6ba388f110487f9d1e63f77b22d86e916b"),
			Price:   fptr(0.065445),
			Adid:    sptr("529833ce55314b19e8796116"),
			Nurl:    sptr("http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}"),
			Adm:     sptr("<iframe src=\"foo.bar\"/>"),
			Adomain: []string{},
			Attr:    []int{},
			Cid:     sptr("529833ce55314b19e8796116"),
			Crid:    sptr("529833ce55314b19e8796116_1385706446"),
		}

		Expect(resp.Seatbid[0].Bid[0]).To(Equal(bid))
	})
})
