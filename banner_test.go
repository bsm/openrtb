package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBanner_IsTopFrame(t *testing.T) {
	b := &Banner{}
	assert.Equal(t, b.IsTopFrame(), false)
}

func TestBanner_Position(t *testing.T) {
	b := &Banner{}
	assert.Equal(t, b.Position(), 0)
}

func TestBanner_WithDefaults(t *testing.T) {
	b := &Banner{}
	banner := b.WithDefaults()
	assert.Equal(t, *banner.Topframe, 0)
	assert.Equal(t, *banner.Pos, AD_POS_UNKNOWN)
}
