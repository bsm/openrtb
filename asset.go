package openrtb

// TODO: Add description, use cases and other details
type Asset struct {
	Id       *int         `json:"id"`                 // Unique asset ID, assigned by exchange
	Required *int         `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *Title       `json:"title,omitempty"`    // Title object for title assets
	Img      *Image       `json:"img,omitempty"`      // Image object for image assets
	Video    *NativeVideo `json:"video,omitempty"`    // Video object for video assets
	Data     *NativeData  `json:"data,omitempty"`     // Data object for ratings, price, etc.
	Ext      Extensions   `json:"ext,omitempty"`
}

type Title struct {
	Len *int       `json:"len"` // Maximum length of the text in the title element
	Ext Extensions `json:"ext,omitempty"`
}

type Image struct {
	Type  *int       `json:"type,omitempty"`  // Type ID of the image element supported by the publisher
	W     *int       `json:"w,omitempty"`     // Width
	H     *int       `json:"h,omitempty"`     // Height
	Wmax  *int       `json:"wmax,omitempty"`  // Width maximum
	Hmax  *int       `json:"hmax,omitempty"`  // Height maximum
	Wmin  *int       `json:"wmin,omitempty"`  // Width minimum
	Hmin  *int       `json:"hmin,omitempty"`  // Height minimum
	Mimes []string   `json:"mimes,omitempty"` // Whitelist of content MIME types supported
	Ext   Extensions `json:"ext,omitempty"`
}

type NativeData struct {
	Type *int       `json:"type"`          // Type ID of the image element supported by the publisher
	Len  *int       `json:"len,omitempty"` // Maximum length of the text in the element’s response
	Ext  Extensions `json:"ext,omitempty"`
}

// Applies defaults
func (a *Asset) WithDefaults() *Asset {
	if a.Required == nil {
		a.Required = new(int)
		*a.Required = 0
	}
	return a
}
