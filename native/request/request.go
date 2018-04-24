package request

import "github.com/tisonet/openrtb-easyjson"

type LayoutID int

const (
	LayoutContentWall          LayoutID = 1
	LayoutAppWall              LayoutID = 2
	LayoutNewsFeed             LayoutID = 3
	LayoutChatList             LayoutID = 4
	LayoutCarousel             LayoutID = 5
	LayoutContentStream        LayoutID = 6
	LayoutGridAdjoiningContent LayoutID = 7
)

type AdUnitID int

const (
	AdUnitPaidSearch            AdUnitID = 1
	AdUnitRecommendationWidget  AdUnitID = 2
	AdUnitPromotedListings      AdUnitID = 3
	AdUnitInAdWithNativeElement AdUnitID = 4
	AdUnitCustom                AdUnitID = 5
)

type ContextTypeID int

const (
	ContextTypeContent ContextTypeID = 1 // newsfeed, article, image gallery, video gallery
	ContextTypeSocial  ContextTypeID = 2 // social network feed, email, chat
	ContextTypeProduct ContextTypeID = 3 // product listings, details, recommendations, reviews
)

type ContextSubTypeID int

const (
	ContextSubTypeGeneral        ContextSubTypeID = 10
	ContextSubTypeArticle        ContextSubTypeID = 11
	ContextSubTypeVideo          ContextSubTypeID = 12
	ContextSubTypeAudio          ContextSubTypeID = 13
	ContextSubTypeImage          ContextSubTypeID = 14
	ContextSubTypeUserGenerated  ContextSubTypeID = 15
	ContextSubTypeSocial         ContextSubTypeID = 20
	ContextSubTypeEmail          ContextSubTypeID = 21
	ContextSubTypeChat           ContextSubTypeID = 22
	ContextSubTypeSelling        ContextSubTypeID = 30
	ContextSubTypeAppStore       ContextSubTypeID = 31
	ContextSubTypeProductReviews ContextSubTypeID = 32
)

type PlacementTypeID int

const (
	PlacementTypeInFeed         PlacementTypeID = 1 // In the feed of content - for example as an item inside the organic feed/grid/listing/carousel
	PlacementTypeAtomic         PlacementTypeID = 2 // In the atomic unit of the content - IE in the article page or single image page
	PlacementTypeOutside        PlacementTypeID = 3 // Outside the core content - for example in the ads section on the right rail, as a banner-style placement near the content, etc.
	PlacementTypeRecommendation PlacementTypeID = 4 // Recommendation widget, most commonly presented below the article content
)

// This object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Request struct {
	Ver              string            `json:"ver,omitempty"`            // Version of the Native Markup
	LayoutID         LayoutID          `json:"layout,omitempty"`         // DEPRECATED The Layout ID of the native ad
	AdUnitID         AdUnitID          `json:"adunit,omitempty"`         // DEPRECATED The Ad unit ID of the native ad
	ContextTypeID    ContextTypeID     `json:"context,omitempty"`        // The context in which the ad appears
	ContextSubTypeID ContextSubTypeID  `json:"contextsubtype,omitempty"` // A more detailed context in which the ad appears
	PlacementTypeID  PlacementTypeID   `json:"plcmttype,omitempty"`      // The design/format/layout of the ad unit being offered
	PlacementCount   int               `json:"plcmtcnt,omitempty"`       // The number of identical placements in this Layout
	Sequence         int               `json:"seq,omitempty"`            // 0 for the first ad, 1 for the second ad, and so on
	Assets           []Asset           `json:"assets"`                   // An array of Asset Objects
	Ext              openrtb.Extension `json:"ext,omitempty"`
}
