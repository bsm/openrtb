package openrtb

import "github.com/goccy/go-json"

// Device object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	UA           string          `json:"ua,omitempty"`             // User agent
	Sua          *UserAgent      `json:"sua,omitempty"`            // Structured User agent. It's more accurate than UA
	Geo          *Geo            `json:"geo,omitempty"`            // Location of the device assumed to be the user’s current location
	DNT          int             `json:"dnt,omitempty"`            // "1": Do not track
	LMT          int             `json:"lmt,omitempty"`            // "1": Limit Ad Tracking
	IP           string          `json:"ip,omitempty"`             // IPv4
	IPv6         string          `json:"ipv6,omitempty"`           // IPv6
	DeviceType   DeviceType      `json:"devicetype,omitempty"`     // The general type of device.
	Make         string          `json:"make,omitempty"`           // Device make
	Model        string          `json:"model,omitempty"`          // Device model
	OS           string          `json:"os,omitempty"`             // Device OS
	OSVersion    string          `json:"osv,omitempty"`            // Device OS version
	HWVersion    string          `json:"hwv,omitempty"`            // Hardware version of the device (e.g., "5S" for iPhone 5S).
	Height       int             `json:"h,omitempty"`              // Physical height of the screen in pixels.
	Width        int             `json:"w,omitempty"`              // Physical width of the screen in pixels.
	PPI          int             `json:"ppi,omitempty"`            // Screen size as pixels per linear inch.
	PixelRatio   float64         `json:"pxratio,omitempty"`        // The ratio of physical pixels to device independent pixels.
	JS           int             `json:"js,omitempty"`             // Javascript status ("0": Disabled, "1": Enabled)
	GeoFetch     int             `json:"geofetch,omitempty"`       // Indicates if the geolocation API will be available to JavaScript code running in the banner,
	FlashVersion string          `json:"flashver,omitempty"`       // Flash version
	Language     string          `json:"language,omitempty"`       // Browser language using ISO-639-1-alpha-2. Only one of language or langb should be present.
	LanguageB    string          `json:"langb,omitempty"`          // Browser language using IETF BCP 47. Only one of language or langb should be present.
	Carrier      string          `json:"carrier,omitempty"`        // Carrier or ISP derived from the IP address
	MCCMNC       string          `json:"mccmnc,omitempty"`         // Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA).
	ConnType     ConnType        `json:"connectiontype,omitempty"` // Network connection type.
	IFA          string          `json:"ifa,omitempty"`            // Native identifier for advertisers
	IDSHA1       string          `json:"didsha1,omitempty"`        // SHA1 hashed device ID
	IDMD5        string          `json:"didmd5,omitempty"`         // MD5 hashed device ID
	PIDSHA1      string          `json:"dpidsha1,omitempty"`       // SHA1 hashed platform device ID
	PIDMD5       string          `json:"dpidmd5,omitempty"`        // MD5 hashed platform device ID
	MacSHA1      string          `json:"macsha1,omitempty"`        // SHA1 hashed device ID; IMEI when available, else MEID or ESN
	MacMD5       string          `json:"macmd5,omitempty"`         // MD5 hashed device ID; IMEI when available, else MEID or ESN
	Ext          json.RawMessage `json:"ext,omitempty"`
}
