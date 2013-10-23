package openrtb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSite_IsPrivacyPolicy(t *testing.T) {
	s := &Site{}
	assert.Equal(t, s.IsPrivacyPolicy(), false)
}

func TestSite_WithDefaults(t *testing.T) {
	assert.Equal(t, "pending", "TODO")
}
