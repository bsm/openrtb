package request

import "encoding/json"

type Image_Type int

const (
	Image_Icon Image_Type = 1 // Icon image
	Image_Logo Image_Type = 2 // Logo image for the brand/app
	Image_Main Image_Type = 3 // Large image preview for the ad
)

type Image struct {
	Type Image_Type `json:"type,omitempty"` // Type ID of the image element supported by the publisher

	Width     int `json:"w,omitempty"`    // Width of the image in pixels
	WidthMin  int `json:"wmin,omitempty"` // The minimum requested width of the image in pixels
	Height    int `json:"h,omitempty"`    // Height of the image in pixels
	HeightMin int `json:"hmin,omitempty"` // The minimum requested height of the image in pixels
	// Either h/w or hmin/wmin should be transmitted. If only h/w is included, it
	// should be considered an exact requirement
	Mimes []string        `json:"mimes,omitempty"` // Whitelist of content MIME types supported
	Ext   json.RawMessage `json:"ext,omitempty"`
}
