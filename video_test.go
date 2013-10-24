package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVideo_Seq(t *testing.T) {
	v := &Video{}
	assert.Equal(t, v.Seq(), 1)
}

func TestVideo_IsBoxingAllowed(t *testing.T) {
	v := &Video{}
	assert.Equal(t, v.IsBoxingAllowed(), true)
}

func TestVideo_Position(t *testing.T) {
	v := &Video{}
	assert.Equal(t, v.Position(), 0)
}

func TestVideo_Valid(t *testing.T) {
	v := &Video{}

	// blank Video
	ok, err := v.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: video has no mimes")
	}

	// with Mimes
	v.Mimes = []string{"RAND_KEY"}
	ok, err = v.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: video linearity missing")
	}

	// with Linearity
	v.SetLinearity(2)
	ok, err = v.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: video minduration missing")
	}

	// with Maxduration
	v.SetMinduration(1)
	ok, err = v.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: video maxduration missing")
	}

	// with Maxduration
	v.SetMaxduration(5)
	ok, err = v.Valid()
	assert.Equal(t, ok, false)
	if err != nil {
		assert.Equal(t, err.Error(), "openrtb parse: video protocol missing")
	}

	// with Protocol
	v.SetProtocol(1)
	ok, err = v.Valid()
	assert.Equal(t, ok, true)
}

func TestVideo_WithDefaults(t *testing.T) {
	v := &Video{}
	vid := v.WithDefaults()

	assert.Equal(t, *vid.Sequence, 1)
	assert.Equal(t, *vid.Boxingallowed, 1)
	assert.Equal(t, *vid.Pos, AD_POS_UNKNOWN)
}
