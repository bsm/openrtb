package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb/v3"
)

func TestDevice(t *testing.T) {
	var subject *Device
	if err := fixture("device", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Device{
		DNT: 0,
		UA:  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		IP:  "123.145.167.189",
		Geo: &Geo{
			Lat:     35.012345,
			Lon:     -115.12345,
			Country: "USA",
			Region:  "CA",
			Metro:   "803",
			City:    "Los Angeles",
			Zip:     "90049",
		},
		PIDSHA1:    "AA000DFE74168477C70D291f574D344790E0BB11",
		PIDMD5:     "AA003EABFB29E6F759F3BDAB34E50BB11",
		Carrier:    "310-410",
		Language:   "en",
		Make:       "Apple",
		Model:      "iPhone",
		OS:         "iOS",
		OSVer:      "6.1",
		JS:         1,
		ConnType:   ConnTypeCell,
		MCCMNC:     "722-341",
		DeviceType: DeviceTypeMobile,
		OAID:       "1fe9a970-efbb-29e0-0bdd-f5dbbf751ab5",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
