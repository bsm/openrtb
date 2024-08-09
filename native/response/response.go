package response

import (
	"encoding/json"

	"github.com/octoclick/openrtb-easyjson"
)

// Response is the native object is the top level JSON object which identifies a native response.
type Response struct {
	Version       openrtb.StringOrNumber `json:"ver,omitempty"`           // Version of the Native Markup
	Assets        []Asset                `json:"assets"`                  // An array of Asset Objects
	Link          Link                   `json:"link"`                    // Destination Link. This is default link object for the ad
	ImpTrackers   []string               `json:"imptrackers,omitempty"`   // Array of impression tracking URLs, expected to return a 1x1 image or 204 response
	JSTracker     string                 `json:"jstracker,omitempty"`     // Optional JavaScript impression tracker. This is a valid HTML, Javascript is already wrapped in <script> tags. It should be executed at impression time where it can be supported
	EventTrackers []EventTrackers        `json:"eventtrackers,omitempty"` // Array of tracking objects to run with the ad, in response to the declared supported methods in the request. Replaces imptrackers and jstracker, to be deprecated.
	Ext           json.RawMessage        `json:"ext,omitempty"`
}
