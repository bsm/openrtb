package openrtb

import (
	"encoding/json"
	"io"
)

type NativeAdm struct {
	Native *NativeResponse
}

type NativeResponse struct {
	Ver         *string         `json:"ver,omitempty"` // Version of the Native Markup version in use
	Assets      []ResponseAsset `json:"assets"`        // Array of Asset Objects
	Link        *Link           `json:"link"`
	Imptrackers []string        `json:"imptrackers,omitempty"`
	Jstracker   *string         `json:"jstracker, omitempty"`
	Ext         Extensions      `json:"ext,omitempty"`
}

type ResponseAsset struct {
	Id       *int           `json:"id"`                 // Unique asset ID, assigned by exchange
	Required *int           `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *ResponseTitle `json:"title,omitempty"`    // Title object for title assets
	Img      *ResponseImg   `json:"img,omitempty"`      // Image object for image assets
	Video    *ResponseVideo `json:"video,omitempty"`    // Video object for video assets
	Data     *ResponseData  `json:"data,omitempty"`     // Data object for ratings, price, etc.
	Link     *Link          `json:"link"`
	Ext      Extensions     `json:"ext,omitempty"`
}

type Link struct {
	Url           *string    `json:"url"`
	Clicktrackers []string   `json:"clicktrackers,omitempty"`
	Fallback      *string    `json:"fallback,omitempty"`
	Ext           Extensions `json:"ext,omitempty"`
}

type ResponseTitle struct {
	Text *string    `json:"text"`
	Ext  Extensions `json:"ext,omitempty"`
}

type ResponseImg struct {
	Url *string    `json:"url"`
	W   *int       `json:"w,omitempty"` // Width
	H   *int       `json:"h,omitempty"` // Height
	Ext Extensions `json:"ext,omitempty"`
}

type ResponseData struct {
	Label *string    `json:"label,omitempty"`
	Value *string    `json:"value"`
	Ext   Extensions `json:"ext,omitempty"`
}

type ResponseVideo struct {
	Vasttag *string `json:"vasttag"`
}

// Set the Asset
func (nativeResponse *NativeResponse) SetAssets(a *ResponseAsset) *NativeResponse {
	nativeResponse.Assets = append(nativeResponse.Assets, *a)
	return nativeResponse
}

//Parses an OpenRTB Native Response from an io.Reader
func ParseNativeAdm(reader io.Reader) (adm *NativeAdm, err error) {
	dec := json.NewDecoder(reader)
	if err = dec.Decode(&adm); err != nil {
		return nil, err
	}
	adm.Native = adm.Native.WithDefaults()
	return adm, nil
}

//Parses an OpenRTB Native Response from bytes
func ParseNativeAdmBytes(data []byte) (adm *NativeAdm, err error) {
	if err = json.Unmarshal(data, &adm); err != nil {
		return nil, err
	}
	adm.Native = adm.Native.WithDefaults()
	return adm, nil
}

// Applies NativeResponse defaults
func (resp *NativeResponse) WithDefaults() *NativeResponse {
	if resp.Ver == nil {
		resp.Ver = new(string)
		*resp.Ver = "1"
	}
	for id, asset := range resp.Assets {
		resp.Assets[id] = *asset.WithDefaults()
	}
	return resp
}

// Applies ResponseAsset defaults
func (asset *ResponseAsset) WithDefaults() *ResponseAsset {
	if asset.Required == nil {
		asset.Required = new(int)
		*asset.Required = 0
	}
	return asset
}
