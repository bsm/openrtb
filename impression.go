package openrtb

// The "imp" object describes the ad position or impression being auctioned.  A single bid request
// can include multiple "imp" objects, a use case for which might be an exchange that supports
// selling all ad positions on a given page as a bundle.  Each "imp" object has a required ID so that
// bids can reference them individually.  An exchange can also conduct private auctions by
// restricting involvement to specific subsets of seats within bidders.
type Impression struct {
	Id                *string // A unique identifier for this impression
	Banner            *Banner
	Video             *Video
	Displaymanager    *string  // Name of ad mediation partner, SDK technology, etc
	Displaymanagerver *string  // Version of the above
	Instl             *int     // Interstitial, Default: 0 ("1": Interstitial, "0": Something else)
	Tagid             *string  // Identifier for specific ad placement or ad tag
	Bidfloor          *float32 // Bid floor for this impression in CPM
	Bidfloorcur       *string  // Currency of bid floor
	Iframebuster      []string // Array of names for supportediframe busters.
	Ext               Extensions
}

// Validates the `imp` object
func (imp *Impression) Valid() (bool, error) {

	if imp.Id == nil {
		return false, errValidationImpId
	} else if imp.Banner != nil && imp.Video != nil {
		return false, errValidationImpBaV
	} else if imp.Video != nil {
		if ok, err := imp.Video.Valid(); !ok {
			return ok, err
		}
	} else {
		return false, errValidationImpBoV
	}

	return true, nil
}

// Returns the `imp` object returning defaults
func (imp *Impression) WithDefaults() *Impression {
	if imp.Instl == nil {
		imp.Instl = new(int)
		*imp.Instl = 0
	}

	if imp.Bidfloor == nil {
		imp.Bidfloor = new(float32)
		*imp.Bidfloor = 0
	}

	if imp.Bidfloorcur == nil {
		imp.Bidfloorcur = new(string)
		*imp.Bidfloorcur = "USD"
	}

	if imp.Banner != nil {
		imp.Banner.WithDefaults()
	}

	if imp.Video != nil {
		imp.Video.WithDefaults()
	}

	return imp
}
