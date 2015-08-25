package openrtb

import "encoding/json"

// The "device" object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	DNT        int             `json:"dnt,omitempty"` // "1": Do not track
	UA         string          `json:"ua,omitempty"`  // User agent
	IP         string          `json:"ip,omitempty"`  // IPv4
	Geo        *Geo            `json:"geo,omitempty"`
	IDSHA1     string          `json:"didSHA1,omitempty"`  // SHA1 hashed device ID
	IDMD5      string          `json:"didMD5,omitempty"`   // MD5 hashed device ID
	PIDSHA1    string          `json:"dpidSHA1,omitempty"` // SHA1 hashed platform device ID
	PIDMD5     string          `json:"dpidMD5,omitempty"`  // MD5 hashed platform device ID
	MacSHA1    string          `json:"macSHA1,omitempty"`  // SHA1 hashed device ID; IMEI when available, else MEID or ESN
	MacMD5     string          `json:"macMD5,omitempty"`   // MD5 hashed device ID; IMEI when available, else MEID or ESN
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
}
