package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpression_Valid(t *testing.T) {
	imp := &Impression{}
	b := &Banner{}
	v := &Video{Mimes: []string{"MIME_123"}}

	// blank Impression
	ok, err := imp.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: impression ID missing")
	}

	// Impression with ID
	imp.SetId("CODE_12")
	ok, err = imp.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: impression has neither a banner nor video")
	}

	// Impression with Banner
	imp.SetBanner(*b)
	ok, err = imp.Valid()
	assert.Equal(t, ok, true)

	// Impression with both Banner and Video
	v.SetLinearity(1).SetMinduration(1).SetMaxduration(5).SetProtocol(1)
	imp.SetVideo(*v)
	ok, err = imp.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: impression has banner and video")
	}

	// Impression with valid attrs
	imp.Banner = nil
	ok, err = imp.Valid()
	assert.Equal(t, ok, true)
}

func TestImpression_WithDefaults(t *testing.T) {
	i := &Impression{}
	b := &Banner{}
	v := &Video{}

	i.SetBanner(*b).SetVideo(*v)
	imp := i.WithDefaults()

	assert.Equal(t, *imp.Instl, 0)
	assert.Equal(t, *imp.Bidfloor, 0)
	assert.Equal(t, *imp.Bidfloorcur, "USD")
	assert.Equal(t, *imp.Banner.Topframe, 0)
	assert.Equal(t, *imp.Banner.Pos, AD_POS_UNKNOWN)
	assert.Equal(t, *imp.Video.Sequence, 1)
	assert.Equal(t, *imp.Video.Boxingallowed, 1)
	assert.Equal(t, *imp.Video.Pos, AD_POS_UNKNOWN)
}
