package response

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	It("should parse correctly", func() {
		Expect(fixture("testdata/response1.json")).To(Equal(&Response{
			Version:     "1.1",
			ImpTrackers: []string{"http://imptracker.com"},
			JSTracker:   "<script>track()</script>",
			Link: Link{
				URL: "http://i.am.a/URL",
			},
			Assets: []Asset{
				{ID: 123, Required: 1, Title: &Title{Text: "Learn about this awesome thing"}},
				{ID: 124, Required: 1, Image: &Image{URL: "http://www.myads.com/thumbnail1.png"}},
				{ID: 128, Required: 1, Image: &Image{URL: "http://www.myads.com/largethumb1.png"}},
				{ID: 126, Required: 1, Data: &Data{Value: "My Brand"}},
				{ID: 127, Required: 1, Data: &Data{Value: "Learn all about this awesome story of someone using my product."}},
				{ID: 4, Video: &Video{VASTTag: "<VAST version=\"2.0\"></VAST>"}},
				{ID: 5, Link: &Link{URL: "http://landing.com", ClickTrackers: []string{"http://tracker.com"}, FallbackURL: "http://fallback.com"}},
			},
		}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openrtb/native/response")
}

func fixture(path string) (res *Response) {
	enc, err := ioutil.ReadFile(path)
	Expect(err).ToNot(HaveOccurred())
	Expect(json.Unmarshal(enc, &res)).To(Succeed())
	return res
}
