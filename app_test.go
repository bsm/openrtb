package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_IsPrivacyPolicy(t *testing.T) {
	a := &App{}
	assert.Equal(t, a.IsPrivacyPolicy(), false)
}

func TestApp_IsPaid(t *testing.T) {
	a := &App{}
	assert.Equal(t, a.IsPaid(), false)
}

func TestApp_WithDefaults(t *testing.T) {
	assert.Equal(t, "pending", "TODO")
}
