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

func TestSite_IsPrivacyPolicy(t *testing.T) {
	s := &Site{}
	assert.Equal(t, s.IsPrivacyPolicy(), false)
}

func TestApp_IsPrivacyPolicy(t *testing.T) {
	a := &App{}
	assert.Equal(t, a.IsPrivacyPolicy(), false)
}

func TestApp_IsPaid(t *testing.T) {
	a := &App{}
	assert.Equal(t, a.IsPaid(), false)
}

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
