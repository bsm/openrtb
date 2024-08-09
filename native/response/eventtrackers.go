package response

import "encoding/json"

// EventTrackers object contains response data.
type EventTrackers struct {
	Event      EventType           `json:"event"`
	Method     EventTrackingMethod `json:"method"`
	Url        string              `json:"url,omitempty"`        //The URL of the image or js. Required for image or js, optional for custom.
	CustomData map[string]any      `json:"customdata,omitempty"` //To be agreed individually with the exchange, an array of key:value objects for custom tracking, for example the account number of the DSP with a tracking company. IE {“accountnumber”:”123”}
	Ext        json.RawMessage     `json:"ext,omitempty"`
}

// EventType ...
type EventType int

const (
	ImpressionEvent = iota + 1 // Impression
	ViewableMRC50              // Visible impression using MRC definition at 50% in view for 1 second
	ViewableMRC100             // 100% in view for 1 second (ie GroupMstandard)
	ViewableVideo50            // Visible impression for video using MRC definition at 50% in view for 2 seconds
)

// EventTrackingMethod ...
type EventTrackingMethod int

const (
	ImageMethod      = iota + 1 // Image-pixel tracking - URL provided will be inserted as a 1x1 pixel at the time of the event
	JavaScriptMethod            // Javascript-based tracking - URL provided will be inserted as a js tag at the time of the event.
)
