package response

import "github.com/Upliner/openrtb"

type Image struct {
	URL    string            `json:"url,omitempty"` // URL of the image asset
	Width  int               `json:"w,omitempty"`   // Width of the image in pixels
	Height int               `json:"h,omitempty"`   // Height of the image in pixels
	Ext    openrtb.Extension `json:"ext,omitempty"`
}
