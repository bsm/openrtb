package request

import "github.com/Upliner/openrtb"

// The main container object for each asset requested or supported by Exchange
// on behalf of the rendering client.  Only one of the {title,img,video,data}
// objects should be present in each object.  The id is to be unique within the
// AssetObject array so that the response can be aligned.
type Asset struct {
	ID       int               `json:"id"`                 // Unique asset ID, assigned by exchange
	Required int               `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *Title            `json:"title,omitempty"`    // Title object for title assets
	Image    *Image            `json:"img,omitempty"`      // Image object for image assets
	Video    *Video            `json:"video,omitempty"`    // Video object for video assets
	Data     *Data             `json:"data,omitempty"`     // Data object for brand name, description, ratings, prices etc.
	Ext      openrtb.Extension `json:"ext,omitempty"`
}
