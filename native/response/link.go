package response

import "encoding/json"

type Link struct {
	URL           string          `json:"url"`                     // Landing URL of the clickable link
	ClickTrackers []string        `json:"clicktrackers,omitempty"` // List of third-party tracker URLs to be fired on click of the URL
	FallbackURL   string          `json:"fallback,omitempty"`      // Fallback URL for deeplink. To be used if the URL given in url is not supported by the device.
	Ext           json.RawMessage `json:"ext,omitempty"`
}
