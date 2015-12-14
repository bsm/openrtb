package openrtb

import "encoding/json"

// The "device" object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	DNT        int             `json:"dnt,omitempty"` // "1": Do not track
	LMT        int             `json:"lmt,omitempty"` // "1": Limit Ad Tracking
	UA         string          `json:"ua,omitempty"`  // User agent
	IP         string          `json:"ip,omitempty"`  // IPv4
	Geo        *Geo            `json:"geo,omitempty"`
	IDSHA1     string          `json:"didsha1,omitempty"`  // SHA1 hashed device ID
	IDMD5      string          `json:"didmd5,omitempty"`   // MD5 hashed device ID
	PIDSHA1    string          `json:"dpidsha1,omitempty"` // SHA1 hashed platform device ID
	PIDMD5     string          `json:"dpidmd5,omitempty"`  // MD5 hashed platform device ID
	MacSHA1    string          `json:"macsha1,omitempty"`  // SHA1 hashed device ID; IMEI when available, else MEID or ESN
	MacMD5     string          `json:"macmd5,omitempty"`   // MD5 hashed device ID; IMEI when available, else MEID or ESN
	IPv6       string          `json:"ipv6,omitempty"`     // IPv6
	Carrier    string          `json:"carrier,omitempty"`  // Carrier or ISP derived from the IP address
	Language   string          `json:"language,omitempty"` // Browser language
	Make       string          `json:"make,omitempty"`     // Device make
	Model      string          `json:"model,omitempty"`    // Device model
	OS         string          `json:"os,omitempty"`       // Device OS
	OSVer      string          `json:"osv,omitempty"`      // Device OS version
	JS         int             `json:"js,omitempty"`       // Javascript status ("0": Disabled, "1": Enabled)
	ConnType   int             `json:"connectiontype,omitempty"`
	DeviceType int             `json:"devicetype,omitempty"`
	FlashVer   string          `json:"flashver,omitempty"` // Flash version
	IFA        string          `json:"ifa,omitempty"`      // Native identifier for advertisers
	Ext        json.RawMessage `json:"ext,omitempty"`
	H          int             `json:"h,omitempty"`       // Physical height of the screen in pixels.
	W          int             `json:"w,omitempty"`       // Physical width of the screen in pixels.
	PPI        int             `json:"ppi,omitempty"`     // Screen size as pixels per linear inch.
	PxRatio    float64         `json:"pxratio,omitempty"` // The ratio of physical pixels to device independent pixels.
	HwVer      string          `json:"hwv,omitempty"`     // Hardware version of the device (e.g., "5S" for iPhone 5S).
}
