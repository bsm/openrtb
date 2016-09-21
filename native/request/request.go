package request

import "encoding/json"

type Request_Layout int

const (
	Request_ContentWall          Request_Layout = 1
	Request_AppWall              Request_Layout = 2
	Request_NewsFeed             Request_Layout = 3
	Request_ChatList             Request_Layout = 4
	Request_Carousel             Request_Layout = 5
	Request_ContentStream        Request_Layout = 6
	Request_GridAdjoiningContent Request_Layout = 7
)

type Request_AdUnit int

const (
	Request_PaidSearchUnits            Request_AdUnit = 1
	Request_RecommendationWidgets      Request_AdUnit = 2
	Request_PromotedListings           Request_AdUnit = 3
	Request_InAdWithNativeElementUnits Request_AdUnit = 4
	Request_Custom                     Request_AdUnit = 5
)

type Request_Context int

const (
	Request_Content Request_Context = 1 // newsfeed, article, image gallery, video gallery
	Request_Social  Request_Context = 2 // social network feed, email, chat
	Request_Product Request_Context = 3 // product listings, details, recommendations, reviews
)

type Request_ContextSubType int

const (
	Request_General        Request_ContextSubType = 10
	Request_Article        Request_ContextSubType = 11
	Request_Video          Request_ContextSubType = 12
	Request_Audio          Request_ContextSubType = 13
	Request_Image          Request_ContextSubType = 14
	Request_UserGenerated  Request_ContextSubType = 15
	Request_SubSocial      Request_ContextSubType = 20
	Request_Email          Request_ContextSubType = 21
	Request_Chat           Request_ContextSubType = 22
	Request_Selling        Request_ContextSubType = 30
	Request_AppStore       Request_ContextSubType = 31
	Request_ProductReviews Request_ContextSubType = 32
)

type Request_PlacementType int

const (
	Request_InFeed         Request_PlacementType = 1 // In the feed of content - for example as an item inside the organic feed/grid/listing/carousel
	Request_Atomic         Request_PlacementType = 2 // In the atomic unit of the content - IE in the article page or single image page
	Request_Outside        Request_PlacementType = 3 // Outside the core content - for example in the ads section on the right rail, as a banner-style placement near the content, etc.
	Request_Recommendation Request_PlacementType = 4 // Recommendation widget, most commonly presented below the article content
)

// This object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Request struct {
	Ver            string                 `json:"ver,omitempty"`            // Version of the Native Markup
	Layout         Request_Layout         `json:"layout,omitempty"`         // DEPRECATED The Layout ID of the native ad
	AdUnit         Request_AdUnit         `json:"adunit,omitempty"`         // DEPRECATED The Ad unit ID of the native ad
	Context        Request_Context        `json:"context,omitempty`         // The context in which the ad appears
	ContextSubType Request_ContextSubType `json:"contextsubtype,omitempty"` // A more detailed context in which the ad appears
	PlacementType  Request_PlacementType  `json:"plcmttype,omitempty"`      // The design/format/layout of the ad unit being offered
	PlacementCount int                    `json:"plcmtcnt,omitempty"`       // The number of identical placements in this Layout
	Sequence       int                    `json:"seq,omitempty"`            // 0 for the first ad, 1 for the second ad, and so on
	Assets         []Asset                `json:"assets"`                   // An array of Asset Objects
	Ext            json.RawMessage        `json:"ext,omitempty"`
}
