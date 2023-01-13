package openrtb

import "encoding/json"

type UserAgent struct {
	Browsers     BrandVersion    `json:"browsers,omitempty"`     // A browser or similar software component
	Platform     BrandVersion    `json:"platform,omitempty"`     // The user agent’s execution platform / OS
	Mobile       int             `json:"mobile,omitempty"`       // 1 if the agent prefers a "mobile" version of the content, if available, i.e. optimized for small screens or touch input. 0 if the agent prefers the "desktop". Taken from Sec-CH-UAMobile header
	Architecture string          `json:"architecture,omitempty"` // Device’s major binary architecture, e.g. "x86" or "arm". Taken from the Sec-CH-UA-Arch header
	Bitness      string          `json:"bitness,omitempty"`      // Device’s bitness, e.g. "64" for 64-bit architecture. Taken from the Sec-CH-UA-Bitness header
	Model        string          `json:"model,omitempty"`        // Device model. Taken from the Sec-CH-UAModel header
	Source       string          `json:"source,omitempty"`       // The source of data used to create this object, List: User-Agent Source in AdCOM 1.0
	Ext          json.RawMessage `json:"ext,omitempty"`
}

type BrandVersion struct {
	Brand   string          `json:"brand,omitempty"`   // A brand identifier, for example, "Chrome" or "Windows". Taken from the Sec-CH-UA-Full-Version or Sec-CH-UA-Platform header
	Version []string        `json:"version,omitempty"` // A sequence of version components, in descending hierarchical order (major, minor, patch)
	Ext     json.RawMessage `json:"ext,omitempty"`
}
