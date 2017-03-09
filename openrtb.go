package openrtb

// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText = 1
	BannerTypeXHTML     = 2
	BannerTypeJS        = 3
	BannerTypeFrame     = 4
)

// 5.3 Creative Attributes
const (
	CreativeAttributeAudioAdAutoPlay                 = 1
	CreativeAttributeAudioAdUserInitiated            = 2
	CreativeAttributeExpandableAuto                  = 3
	CreativeAttributeExpandableUserInitiatedClick    = 4
	CreativeAttributeExpandableUserInitiatedRollover = 5
	CreativeAttributeInBannerVideoAdAutoPlay         = 6
	CreativeAttributeInBannerVideoAdUserInitiated    = 7
	CreativeAttributePop                             = 8
	CreativeAttributeProvocativeOrSuggestiveImagery  = 9
	CreativeAttributeExtremeAnimation                = 10
	CreativeAttributeSurveys                         = 11
	CreativeAttributeTextOnly                        = 12
	CreativeAttributeUserInitiated                   = 13
	CreativeAttributeWindowsDialogOrAlert            = 14
	CreativeAttributeHasAudioWithPlayer              = 15
	CreativeAttributeAdProvidesSkipButton            = 16
	CreativeAttributeAdobeFlash                      = 17
)

// 5.4 Ad Position
const (
	AdPosUnknown    = 0
	AdPosAboveFold  = 1
	AdPosBelowFold  = 3
	AdPosHeader     = 4
	AdPosFooter     = 5
	AdPosSidebar    = 6
	AdPosFullscreen = 7
)

// 5.5 Expandable Direction
const (
	ExpDirLeft       = 1
	ExpDirRight      = 2
	ExpDirUp         = 3
	ExpDirDown       = 4
	ExpDirFullScreen = 5
)

// 5.6 API Frameworks
const (
	APIFrameworkVPAID1 = 1
	APIFrameworkVPAID2 = 2
	APIFrameworkMRAID1 = 3
	APIFrameworkORMMA  = 4
	APIFrameworkMRAID2 = 5
)

// 5.7 Video Linearity
const (
	VideoLinearityLinear    = 1
	VideoLinearityNonLinear = 2
)

// 5.8 Video and Audio Bid Response Protocols
const (
	VideoProtoVAST1            = 1
	VideoProtoVAST2            = 2
	VideoProtoVAST3            = 3
	VideoProtoVAST1Wrapper     = 4
	VideoProtoVAST2Wrapper     = 5
	VideoProtoVAST3Wrapper     = 6
	VideoProtoVAST4            = 7
	VideoProtoVAST4Wrapper     = 8
	AudioProtocolDAAST1        = 9
	AudioProtocolDAAST1Wrapper = 10
)

// 5.9 Video Playback Methods
const (
	VideoPlaybackAutoSoundOn  = 1
	VideoPlaybackAutoSoundOff = 2
	VideoPlaybackClickToPlay  = 3
	VideoPlaybackMouseOver    = 4
)

// 5.9 Video Placement Types (Spec 2.5)
const (
	VideoPlacementInStream     = 1
	VideoPlacementInBanner     = 2
	VideoPlacementInArticle    = 3
	VideoPlacementInFeed       = 4
	VideoPlacementInterstitial = 5
)

// 5.10 Video Start Delay
const (
	VideoStartDelayPreRoll         = 0
	VideoStartDelayGenericMidRoll  = -1
	VideoStartDelayGenericPostRoll = -2
)

// 5.11 Video Quality
const (
	VideoQualityUnknown      = 0
	VideoQualityProfessional = 1
	VideoQualityProsumer     = 2
	VideoQualityUGC          = 3
)

// 5.12 VAST Companion Types
const (
	VASTCompanionStatic = 1
	VASTCompanionHTML   = 2
	VASTCompanionIFrame = 3
)

// 5.13 Content Delivery Methods
const (
	ContentDeliveryStreaming   = 1
	ContentDeliveryProgressive = 2
	ContentDeliveryDownload    = 3
)

// 5.14 Content Context
const (
	ContextVideo       = 1
	ContextGame        = 2
	ContextMusic       = 3
	ContextApplication = 4
	ContextText        = 5
	ContextOther       = 6
	ContextUnknown     = 7
)

// 5.15 QAG Media Ratings
const (
	QAGAll    = 1
	QAGOver12 = 2
	QAGMature = 3
)

// 5.16 Location Type
const (
	LocationTypeGPS  = 1
	LocationTypeIP   = 2
	LocationTypeUser = 3
)

// 5.17 Device Type
const (
	DeviceTypeUnknown   = 0
	DeviceTypeMobile    = 1
	DeviceTypePC        = 2
	DeviceTypeTV        = 3
	DeviceTypePhone     = 4
	DeviceTypeTablet    = 5
	DeviceTypeConnected = 6
	DeviceTypeSetTopBox = 7
)

// 5.18 Connection Type
const (
	ConnTypeUnknown  = 0
	ConnTypeEthernet = 1
	ConnTypeWIFI     = 2
	ConnTypeCell     = 3
	ConnTypeCell2G   = 4
	ConnTypeCell3G   = 5
	ConnTypeCell4G   = 6
)

// 5.19 No-Bid Reason Codes
const (
	NBRUnknownError      = 0
	NBRTechnicalError    = 1
	NBRInvalidRequest    = 2
	NBRKnownSpider       = 3
	NBRSuspectedNonHuman = 4
	NBRProxyIP           = 5
	NBRUnsupportedDevice = 6
	NBRBlockedSite       = 7
	NBRUnmatchedUser     = 8
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
	ID     string    `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Cat    []string  `json:"cat,omitempty"` // Array of IAB content categories
	Domain string    `json:"domain,omitempty"`
	Ext    Extension `json:"ext,omitempty"`
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
	Lat           float64   `json:"lat,omitempty"`           // Latitude from -90 to 90
	Lon           float64   `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          int       `json:"type,omitempty"`          // Indicate the source of the geo data
	Accuracy      int       `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int       `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int       `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string    `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string    `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string    `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string    `json:"metro,omitempty"`
	City          string    `json:"city,omitempty"`
	Zip           string    `json:"zip,omitempty"`
	UTCOffset     int       `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           Extension `json:"ext,omitempty"`
}

// This object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID         string    `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerID    string    `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID   string    `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	YOB        int       `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender     string    `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords   string    `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData string    `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo        *Geo      `json:"geo,omitempty"`
	Data       []Data    `json:"data,omitempty"`
	Ext        Extension `json:"ext,omitempty"`
}

// The data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Segment []Segment `json:"segment,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Value string    `json:"value,omitempty"`
	Ext   Extension `json:"ext,omitempty"`
}

// This object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	Coppa int       `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Ext   Extension `json:"ext,omitempty"`
}

// This object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	W   int       `json:"w,omitempty"` // Width in device independent pixels (DIPS).
	H   int       `json:"h,omitempty"` //Height in device independent pixels (DIPS).
	Ext Extension `json:"ext,omitempty"`
}
