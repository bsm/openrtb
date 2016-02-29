package openrtb

import "encoding/json"

// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText int = iota + 1
	BannerTypeXHTML
	BannerTypeJS
	BannerTypeFrame
)

// 5.3 Creative Attributes
// TODO

// 5.4 Ad Position
const (
	AdPosUnknown int = iota
	AdPosAboveFold
	AdPosBelowFold
	AdPosHeader
	AdPosFooter
	AdPosSidebar
	AdPosFullscreen
)

// 5.5 Expandable Direction
const (
	ExpDirLeft int = iota + 1
	ExpDirRight
	ExpDirUp
	ExpDirDown
	ExpDirFullScreen
)

// 5.6 API Frameworks
const (
	APIFrameworkVPAID1 int = iota + 1
	APIFrameworkVPAID2
	APIFrameworkMRAID1
	APIFrameworkORMMA
	APIFrameworkMRAID2
)

// 5.7 Video Linearity
const (
	VideoLinearityLinear int = iota + 1
	VideoLinearityNonLinear
)

// 5.8 Video Bid Response Protocols
const (
	VideoProtoVAST1 int = iota + 1
	VideoProtoVAST2
	VideoProtoVAST3
	VideoProtoVAST1Wrapper
	VideoProtoVAST2Wrapper
	VideoProtoVAST3Wrapper
)

// 5.9 Video Playback Methods
const (
	VideoPlaybackAutoSoundOn int = iota + 1
	VideoPlaybackAutoSoundOff
	VideoPlaybackClickToPlay
	VideoPlaybackMouseOver
)

// 5.10 Video Start Delay
const (
	VideoStartDelayPreRoll         = 0
	VideoStartDelayGenericMidRoll  = -1
	VideoStartDelayGenericPostRoll = -2
)

// 5.11 Video Quality
const (
	VideoQualityUnknown int = iota
	VideoQualityProfessional
	VideoQualityProsumer
	VideoQualityUGC
)

// 5.12 VAST Companion Types
const (
	VASTCompanionStatic int = iota + 1
	VASTCompanionHTML
	VASTCompanionIFrame
)

// 5.13 Content Delivery Methods
const (
	ContentDeliveryStreaming int = iota + 1
	ContentDeliveryProgressive
)

// 5.14 Content Context
const (
	ContextVideo int = iota + 1
	ContextGame
	ContextMusic
	ContextApplication
	ContextText
	ContextOther
	ContextUnknown
)

// 5.15 QAG Media Ratings
const (
	QAGAll int = iota + 1
	QAGOver12
	QAGMature
)

// 5.16 Location Type
const (
	LocationTypeGPS int = iota + 1
	LocationTypeIP
	LocationTypeUser
)

// 5.17 Device Type
const (
	DeviceTypeUnknown int = iota
	DeviceTypeMobile
	DeviceTypePC
	DeviceTypeTV
	DeviceTypePhone
	DeviceTypeTablet
	DeviceTypeConnected
	DeviceTypeSetTopBox
)

// 5.18 Connection Type
const (
	ConnTypeUnknown int = iota
	ConnTypeEthernet
	ConnTypeWIFI
	ConnTypeCell
	ConnTypeCell2G
	ConnTypeCell3G
	ConnTypeCell4G
)

// 5.19 No-Bid Reason Codes
const (
	NBRUnknownError int = iota
	NBRTechnicalError
	NBRInvalidRequest
	NBRKnownSpider
	NBRSuspectedNonHuman
	NBRProxyIP
	NBRUnsupportedDevice
	NBRBlockedSite
	NBRUnmatchedUser
)

/*************************************************************************
 * COMMON OBJECT STRUCTS
 *************************************************************************/

// This object may be useful in the situation where syndicated content contains impressions and
// does not necessarily match the publisher's general content.  The exchange might or might not
// have knowledge of the page where the content is running, as a result of the syndication
// method.  (For example, video impressions embedded in an iframe on an unknown web property
// or device.)
// type Content struct {
// }

// Abstract third-party
type ThirdParty struct {
	ID     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Cat    []string        `json:"cat,omitempty"` // Array of IAB content categories
	Domain string          `json:"domain,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// The publisher object itself and all of its parameters are optional, so default values are not
// provided. If an optional parameter is not specified, it should be considered unknown.
type Publisher ThirdParty

// The producer is useful when content where the ad is shown is syndicated, and may appear on a
// completely different publisher. The producer object itself and all of its parameters are optional,
// so default values are not provided. If an optional parameter is not specified, it should be
// considered unknown.
type Producer ThirdParty

// Note that the Geo Object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Lat           float64         `json:"lat,omitempty"`           // Latitude from -90 to 90
	Lon           float64         `json:"lon,omitempty"`           // Longitude from -180 to 180
	Country       string          `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string          `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string          `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string          `json:"metro,omitempty"`
	City          string          `json:"city,omitempty"`
	Zip           string          `json:"zip,omitempty"`
	Type          int             `json:"type,omitempty"`      // Indicate the source of the geo data
	UTCOffset     int             `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           json.RawMessage `json:"ext,omitempty"`
}

// This object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID         string          `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerUID   string          `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid or id is recommended.
	YOB        int             `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender     string          `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords   string          `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData string          `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo        *Geo            `json:"geo,omitempty"`
	Data       []Data          `json:"data,omitempty"`
	Ext        json.RawMessage `json:"ext,omitempty"`
}

// The data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string          `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	Segment []Segment       `json:"segment,omitempty"`
	Ext     json.RawMessage `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Value string          `json:"value,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}

// This object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	Coppa int             `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Ext   json.RawMessage `json:"ext,omitempty"`
}
