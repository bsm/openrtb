package native

import "openrtb"

// The Native Object defines the native advertising opportunity available for
// bid via this bid request. It must be included directly in the impression
// object if the impression offered for auction is a native ad format.
type NativeBidRequest struct {
	Version        string                 `json:"ver,omitempty"`
	Layout         int                    `json:"layout,omitempty"`
	AdUnitID       int                    `json:"adunit,omitempty"`
	PlacementCount int                    `json:"plcmtcnt,omitempty"`
	Sequence       int                    `json:"seq,omitempty"`
	Assets         []*Asset               `json:"assets,omitempty"`
	Ext            map[string]interface{} `json:"ext,omitempty"`
}

// The main container object for each asset requested or supported by Exchange
// on behalf of the rendering client. Any object that is required is to be
// flagged as such. Only one of the {title,img,video,data} objects should be
// present in each object.  All others should be null/absent. The id is to be
// unique within the AssetObject array so that the response can be aligned.
type Asset struct {
	ID       int                    `json:"id"`
	Required int                    `json:"required,omitempty"`
	Title    *Title                 `json:"title,omitempty"`
	Image    *Image                 `json:"img,omitempty"`
	Video    *openrtb.Video         `json:"video,omitempty"`
	Data     *Data                  `json:"data,omitempty"`
	Link     *Link                  `json:"link,omitempty"`
	Ext      map[string]interface{} `json:"ext,omitempty"`
}

// The Title object is to be used for title element of the Native ad.
type Title struct {
	Length int                    `json:"len"`
	Ext    map[string]interface{} `json:"ext,omitempty"`
}

// The Image object to be used for all image elements of the Native ad such as
// Icons, Main Image, etc.
type Image struct {
	Type      int                    `json:"type,omitempty"`
	Width     int                    `json:"w,omitempty"`
	MinWidth  int                    `json:"wmin,omitempty"`
	Height    int                    `json:"h,omitempty"`
	MinHeight int                    `json:"hmin,omitempty"`
	MIME      []string               `json:"mimes,omitempty"`
	Ext       map[string]interface{} `json:"ext,omitempty"`
}

// The video object to be used for all video elements supported in the Native
// Ad. This corresponds to the Video object of OpenRTB 2.3.  Exchange
// implementers can impose their own specific restrictions.  type Video
// struct{}

// The Data Object is to be used for all non-core elements of the native unit
// such as Ratings, Review Count, Stars, Download count, descriptions etc.
type Data struct {
	Type   int                    `json:"type"`
	Length int                    `json:"len,omitempty"`
	Ext    map[string]interface{} `json:"ext,omitempty"`
}

// The native object is the top level JSON object which identifies a native
// response.
type NativeBidResponse struct {
	Version   int                    `json:"ver,omitempty"`
	Assets    []*Asset               `json:"assets"`
	Link      *Link                  `json:"link,omitempty"`
	Trackers  []string               `json:"imptrackers,omitempty"`
	JsTracker string                 `json:"jstracker,omitempty"`
	Ext       map[string]interface{} `json:"ext,omitempty"`
}
