package openrtb

import "encoding/json"

// This object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher's general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {
	ID                 string           `json:"id,omitempty"`                 // ID uniquely identifying the content.
	Episode            int              `json:"episode,omitempty"`            // Episode number (typically applies to video content).
	Title              string           `json:"title,omitempty"`              // Content title.
	Series             string           `json:"series,omitempty"`             // Content series.
	Season             string           `json:"season,omitempty"`             // Content season.
	Producer           *Producer        `json:"producer,omitempty"`           // The producer.
	URL                string           `json:"url,omitempty"`                // URL of the content, for buy-side contextualization or review.
	Cat                []string         `json:"url,omitempty"`                // Array of IAB content categories that describe the content.
	VideoQuality       int              `json:"videoquality,omitempty"`       // Video quality per IAB's classification.
	Context            int              `json:"context,omitempty"`            // Type of content (game, video, text, etc.).
	ContentRating      string           `json:"contentrating,omitempty"`      // Content rating (e.g., MPAA).
	UserRating         string           `json:"userrating,omitempty"`         // User rating of the content (e.g., number of stars, likes, etc.).
	QAGMediaRating     int              `json:"qagmediarating,omitempty"`     // Media rating per QAG guidelines.
	Keywords           string           `json:"keywords,omitempty"`           // Comma separated list of keywords describing the content.
	LiveStream         int              `json:"livestream,omitempty"`         // 0 = not live, 1 = content is live (e.g., stream, live blog).
	SourceRelationship int              `json:"sourcerelationship,omitempty"` // 0 = indirect, 1 = direct.
	Len                int              `json:"len,omitempty"`                // Length of content in seconds; appropriate for video or audio.
	Language           string           `json:"language,omitempty"`           // Content language using ISO-639-1-alpha-2.
	Embeddable         int              `json:"embeddable,omitempty"`         // Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	Ext                *json.RawMessage `json:"ext,omitempty"`
}
