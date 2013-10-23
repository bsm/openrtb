package openrtb

const (
	VERSION = "2.1"

	AD_POS_UNKNOWN    = 0
	AD_POS_ABOVE_FOLD = 1
	AD_POS_BELOW_FOLD = 3
	AD_POS_HEADER     = 4
	AD_POS_FOOTER     = 5
	AD_POS_SIDEBAR    = 6
	AD_POS_FULLSCREEN = 7

	BANNER_TYPE_XHTML_TEXT = 1
	BANNER_TYPE_XHTML      = 2
	BANNER_TYPE_JS         = 3
	BANNER_TYPE_FRAME      = 4

	CONN_TYPE_UNKNOWN  = 0
	CONN_TYPE_ETHERNET = 1
	CONN_TYPE_WIFI     = 2
	CONN_TYPE_CELL     = 3
	CONN_TYPE_CELL_2G  = 4
	CONN_TYPE_CELL_3G  = 5
	CONN_TYPE_CELL_4G  = 6

	LOCATION_TYPE_GPS  = 1
	LOCATION_TYPE_IP   = 2
	LOCATION_TYPE_USER = 3

	DEVICE_TYPE_UNKNOWN = 0
	DEVICE_TYPE_MOBILE  = 1
	DEVICE_TYPE_PC      = 2
	DEVICE_TYPE_TV      = 3

	VIDEO_LINEARITY_LINEAR     = 1
	VIDEO_LINEARITY_NON_LINEAR = 2

	VIDEO_PLAYBACK_AUTO_SOUND_ON  = 1
	VIDEO_PLAYBACK_AUTO_SOUND_OFF = 2
	VIDEO_PLAYBACK_CLICK_TO_PLAY  = 3
	VIDEO_PLAYBACK_MOUSE_OVER     = 4

	VIDEO_START_DELAY_PRE_ROLL  = 0
	VIDEO_START_DELAY_MID_ROLL  = -1
	VIDEO_START_DELAY_POST_ROLL = -2

	VIDEO_QUALITY_UNKNOWN      = 0
	VIDEO_QUALITY_PROFESSIONAL = 1
	VIDEO_QUALITY_PROSUMER     = 2
	VIDEO_QUALITY_UGC          = 3
)

// At least one of Bid is required.
// A bid response can contain multiple “seatbid” objects, each on behalf of a different bidder seat.
// Seatbid object can contain multiple bids each pertaining to a different impression on behalf of a seat.
// Each “bid” object must include the impression ID to which it pertains as well as the bid price.
// Group attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is
// only interested in winning any if it can win them all (i.e., all or nothing).
type Seatbid struct {
	Bid   []Bid      `json:"id"`             // Array of bid objects; each realtes to an imp, if exchange supported can have many bid objects.
	Seat  *string    `json:"seat,omiempty"`  // ID of the bidder seat optional string ID of the bidder seat on whose behalf this bid is made.
	Group *int       `json:"group,omiempty"` // '1' means impression must be won-lost as a group; default is '0'.
	Ext   Extensions `json:"ext,omiempty"`
}

// ID, Impid and Price are required; all other optional.
// If the bidder wins the impression, the exchange calls notice URL (nurl)
// a) to inform the bidder of the win;
// b) to convey certain information using substitution macros.
// Adomain can be used to check advertiser block list compliance.
// Cid can be used to block ads that were previously identified as inappropriate.
// Substitution macros may allow a bidder to use a static notice URL for all of its bids.
type Bid struct {
	Id      *string    `json:"id"`
	Impid   *string    `json:"impid"`             // Required string ID of the impression object to which this bid applies.
	Price   *float32   `json:"price"`             // Bid price in CPM. Suggests using integer math for accounting to avoid rounding errors.
	Adid    *string    `json:"adid,omitempty"`    // References the ad to be served if the bid wins.
	Nurl    *string    `json:"nurl,omitempty"`    // Win notice URL.
	Adm     *string    `json:"adm,omitempty"`     // Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	Adomain []string   `json:"adomain,omitempty"` // Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	Iurl    *string    `json:"iurl,omitempty"`    // Sample image URL.
	Cid     *string    `json:"cid,omitempty"`     // Campaign ID that appears with the Ad markup.
	Crid    *string    `json:"crid,omitempty"`    // Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Attr    []int      `json:"attr,omitempty"`    // Array of creative attributes.
	Ext     Extensions `json:"ext,omitempty"`
}

// This object may be useful in the situation where syndicated content contains impressions and
// does not necessarily match the publisher’s general content.  The exchange might or might not
// have knowledge of the page where the content is running, as a result of the syndication
// method.  (For example, video impressions embedded in an iframe on an unknown web property
// or device.)
// type Content struct {
// }

// Abstract third-party
type ThirdParty struct {
	Id     *string
	Name   *string
	Cat    []string // Array of IAB content categories
	Domain *string
	Ext    Extensions
}

// The publisher object itself and all of its parameters are optional, so default values are not
// provided. If an optional parameter is not specified, it should be considered unknown.
type Publisher ThirdParty

// The producer is useful when content where the ad is shown is syndicated, and may appear on a
// completely different publisher. The producer object itself and all of its parameters are optional,
// so default values are not provided. If an optional parameter is not specified, it should be
// considered unknown.
type Producer ThirdParty

// Note that the Geo Object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Lat           *float32 // Latitude from -90 to 90
	Lon           *float32 // Longitude from -180 to 180
	Country       *string  // Country using ISO 3166-1 Alpha 3
	Region        *string  // Region using ISO 3166-2
	Regionfips104 *string  // Region of a country using fips 10-4
	Metro         *string
	City          *string
	Zip           *string
	Type          *int // Indicate the source of the geo data
	Ext           Extensions
}

// The "user" object contains information known or derived about the human user of the device.
// Note that the user ID is an exchange artifact (refer to the "device" object for hardware or
// platform derived IDs) and may be subject to rotation policies. However, this user ID must be
// stable long enough to serve reasonably as the basis for frequency capping.
type User struct {
	Id         *string // Unique consumer ID of this user on the exchange
	Buyeruid   *string // Buyer's user ID
	Yob        *int    // Year-of-birth
	Gender     *string // Gender ("M": male, "F" female, "O" Other)
	Keywords   *string
	Customdata *string
	Geo        *Geo
	Data       []Data
	Ext        Extensions
}

// The data and segment objects together allow data about the user to be passed to bidders in the
// bid request.  This data may be from multiple sources (e.g., the exchange itself, third party
// providers) as specified by the data object ID field.  A bid request can mix data objects from
// multiple providers.
type Data struct {
	Id      *string
	Name    *string
	Segment []Segment
	Ext     Extensions
}

// The data and segment objects together allow data about the user to be passed to bidders in the
// bid request.  Segment objects convey specific units of information from the provider identified
// in the parent data object.
type Segment struct {
	Id    *string
	Name  *string
	Value *string
	Ext   Extensions
}

// Private Marketplace Object
type Pmp struct {
	Private *int
	Deals   []Deal
	Ext     Extensions
}

// Private Marketplace Deal
type Deal struct {
	Id       *string  // Unique deal ID
	At       *int     // Auction type, Default: 2 ("1": first price auction, "2": then second price auction)
	Seats    []string // List of seat IDs attached to this deal
	Type     *int     // Deal indicator ("1": Eplicit Deal, "2": Trading Agreement Deal)
	Bidfloor *float32
	Ext      Extensions
}

// General Extensions
type Extensions map[string]string
