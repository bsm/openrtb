package openrtb_test

import (
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestContent(t *testing.T) {
	var subject *Content
	if err := fixture("content", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Content{
		ID:    "1",
		Title: "Video",
		Producer: &Producer{
			ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
			Name:   "yahoo",
			Domain: "www.yahoo.com",
		},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func TestContent_quoted(t *testing.T) {
	var subject *Content
	if err := fixture("content.quoted", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Content{
		Keywords:           "Orwell, 1984",
		UserRating:         "3",
		Episode:            1,
		Title:              "This is the video title",
		URL:                "http://cdnp.tremormedia.com/video/1984.flv",
		LiveStream:         0,
		ContentRating:      "G",
		Length:             129,
		Series:             "book reading",
		VideoQuality:       2,
		Context:            1,
		Season:             "1",
		SourceRelationship: 0,
		ID:                 "eb9f13ede5fd225333971523f6042f9d",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
