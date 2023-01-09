package response_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3/native/response"
)

func TestResponse(t *testing.T) {
	var subject *Response
	if err := fixture("testdata/response1.json", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Response{
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
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func fixture(path string, v interface{}) error {
	bin, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(bin, v)
}
