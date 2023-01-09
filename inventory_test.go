package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestApp(t *testing.T) {
	var subject *App
	if err := fixture("app", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &App{
		Inventory: Inventory{
			ID:         "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfIUDA",
			Name:       "Yahoo Weather",
			Categories: []ContentCategory{"weather", ContentCategoryScience, ContentCategoryWeather},
			Publisher: &Publisher{
				ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
				Name:   "yahoo",
				Domain: "www.yahoo.com",
			},
		},
		Bundle:   "628677149",
		Version:  "1.0.2",
		StoreURL: "https://itunes.apple.com/id628677149",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestSite(t *testing.T) {
	var subject *Site
	if err := fixture("site", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Site{
		Inventory: Inventory{
			ID:         "102855",
			Categories: []ContentCategory{ContentCategoryAdvertising},
			Domain:     "http://www.usabarfinder.com",
			Publisher: &Publisher{
				ID:     "8953",
				Name:   "local.com",
				Domain: "local.com",
			},
		},
		Page: "http://eas.usabarfinder.com/eas?cu=13824;cre=mu;target=_blank",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
