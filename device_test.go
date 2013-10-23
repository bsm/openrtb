package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDevice_IsDnt(t *testing.T) {
	d := &Device{}
	assert.Equal(t, d.IsDnt(), false)
}

func TestDevice_IsJs(t *testing.T) {
	d := &Device{}
	assert.Equal(t, d.IsJs(), false)
}

func TestDevice_ConnectionType(t *testing.T) {
	d := &Device{}
	assert.Equal(t, d.ConnectionType(), 0)
}

func TestDevice_DeviceType(t *testing.T) {
	d := &Device{}
	assert.Equal(t, d.DeviceType(), 0)
}

func TestDevice_WithDefaults(t *testing.T) {
	assert.Equal(t, "pending", "TODO")
}
