package openrtb

// The "banner" object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	W        *int     // Width
	H        *int     // Height
	Id       *string  // A unique identifier
	Pos      *int     // Ad Position
	Btype    []int    // Blocked creative types
	Battr    []int    // Blocked creative attributes
	Mimes    []string // Whitelist of content MIME types supported
	Topframe *int     // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	Expdir   []int    // Specify properties for an expandable ad
	Api      []int    // List of supported API frameworks
	Ext      Extensions
}

// Returns topframe status, with default fallback
func (b *Banner) IsTopFrame() bool {
	if b.Topframe != nil {
		return *b.Topframe == 1
	}
	return false
}

// Returns the position, with default fallback
func (b *Banner) Position() int {
	if b.Pos != nil {
		return *b.Pos
	}
	return AD_POS_UNKNOWN
}

// Applies defaults
func (b *Banner) WithDefaults() *Banner {
	if b.Topframe == nil {
		b.Topframe = new(int)
		*b.Topframe = 0
	}
	if b.Pos == nil {
		b.Pos = new(int)
		*b.Pos = AD_POS_UNKNOWN
	}
	return b
}
