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
	d := &Device{}
	device := d.WithDefaults()

	assert.Equal(t, *device.Dnt, 0)
	assert.Equal(t, *device.Js, 0)
	assert.Equal(t, *device.Connectiontype, CONN_TYPE_UNKNOWN)
	assert.Equal(t, *device.Devicetype, DEVICE_TYPE_UNKNOWN)
}
