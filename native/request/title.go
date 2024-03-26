package request

import "github.com/UnityTech/openrtb/v3"

// Title is the native title object.
type Title struct {
	Length int               `json:"len"` // Maximum length of the text in the title element
	Ext    openrtb.Extension `json:"ext,omitempty"`
}
