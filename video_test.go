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
	assert.Equal(t, "pending", "TODO")
}

func TestVideo_WithDefaults(t *testing.T) {
	assert.Equal(t, "pending", "TODO")
}
