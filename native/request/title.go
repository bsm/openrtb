package request

import "github.com/tisonet/openrtb-easyjson"

type Title struct {
	Length int               `json:"len"` // Maximum length of the text in the title element
	Ext    openrtb.Extension `json:"ext,omitempty"`
}
