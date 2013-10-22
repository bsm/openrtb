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

// The top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type Request struct {
	Id      *string // Unique ID of the bid request
	Imp     []Impression
	Site    *Site
	App     *App
	Device  *Device
	User    *User
	At      *int     // Auction type, Default: 2 ("1": first price auction, "2": then second price auction)
	Tmax    *int     // Maximum amount of time in milliseconds to submit a bid
	Wseat   []string // Array of buyer seats allowed to bid on this auction
	Allimps *int     // Flag to indicate whether exchange can verify that all impressions offered represent all of the impressions available in context, Default: 0
	Cur     []string // Array of allowed currencies
	Bcat    []string // Blocked Advertiser Categories.
	Badv    []string // Array of strings of blocked toplevel domains of advertisers
	Pmp     *Pmp
	Ext     Extensions
}

// ID and at least one “seatbid” object is required, which contains a bid on at least one impression.
// Other attributes are optional since an exchange may establish default values.
// No-Bids on all impressions should be indicated as a HTTP 204 response.
// For no-bids on specific impressions, the bidder should omit these from the bid response.
type Response struct {
	Id         *string    `json:"id"`                   // Reflection of the bid request ID for logging purposes
	Seatbid    []Seatbid  `json:"seatbid"`              // Array of seatbid objects
	Bidid      *string    `json:"bidid,omitempty"`      // Optional response tracking ID for bidders
	Cur        *string    `json:"cur,omitempty"`        // Bid currency
	Customdata *string    `json:"customdata,omitempty"` // Encoded user features
	Ext        Extensions `json:"ext,omitempty"`        // Custom specifications in Json
}

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

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string // Content MIME types supported.
	Linearity      *int     // Indicates whether the ad impression is linear or non-linear
	Minduration    *int     // Minimum video ad duration in seconds
	Maxduration    *int     // Maximum video ad duration in seconds
	Protocol       *int     // Video bid response protocols
	W              *int     // Width of the player in pixels
	H              *int     // Height of the player in pixels
	Startdelay     *int     // Indicates the start delay in seconds
	Sequence       *int     // Default: 1
	Battr          []int    // Blocked creative attributes
	Maxextended    *int     // Maximum extended video ad duration
	Minbitrate     *int     // Minimum bit rate in Kbps
	Maxbitrate     *int     // Maximum bit rate in Kbps
	Boxingallowed  *int     // If exchange publisher has rules preventing letter boxing
	Playbackmethod []int    // List of allowed playback methods
	Delivery       []int    // List of supported delivery methods
	Pos            *int     // Ad Position
	Companionad    []Banner
	Api            []int // List of supported API frameworks
	Companiontype  []int
	Ext            Extensions
}

// Returns the sequence number, with default fallback
func (v *Video) Seq() int {
	if v.Sequence != nil {
		return *v.Sequence
	}
	return 1
}

// Returns the boxing permission status, with default fallback
func (v *Video) IsBoxingAllowed() bool {
	if v.Boxingallowed != nil {
		return *v.Boxingallowed == 1
	}
	return true
}

// Returns the position, with default fallback
func (v *Video) Position() int {
	if v.Pos != nil {
		return *v.Pos
	}
	return AD_POS_UNKNOWN
}

// A site object should be included if the ad supported content is part of a website (as opposed to
// an application).  A bid request must not contain both a site object and an app object.
type Site struct {
	Id            *string // Site ID on the exchange
	Name          *string // Site name
	Domain        *string
	Cat           []string // Array of IAB content categories
	Sectioncat    []string // Array of IAB content categories for subsection
	Pagecat       []string // Array of IAB content categories for page
	Page          *string  // URL of the page
	Privacypolicy *int     // Default: 1 ("1": site has a privacy policy)
	Ref           *string  // Referrer URL
	Search        *string  // Search string that caused naviation
	Publisher     *Publisher
	// Content       Content
	Keywords []string
	Ext      Extensions
}

// Returns the privacy policy status, with default fallback
func (s *Site) IsPrivacyPolicy() bool {
	if s.Privacypolicy != nil {
		return *s.Privacypolicy == 1
	}
	return false
}

// An "app" object should be included if the ad supported content is part of a mobile application
// (as opposed to a mobile website).  A bid request must not contain both an "app" object and a
// "site" object.
type App struct {
	Id            *string // App ID on the exchange
	Name          *string // App name
	Domain        *string
	Cat           []string // Array of IAB content categories
	Sectioncat    []string // Array of IAB content categories for subsection
	Pagecat       []string // Array of IAB content categories for page
	Ver           *string  // App version
	Bundle        *string  // App bundle or package name
	Privacypolicy *int     // Default: 1 ("1": site has a privacy policy)
	Paid          *int     // "1": Paid, "2": Free
	Publisher     *Publisher
	// Content       Content
	Keywords []string
	Storeurl *string // App store URL for an installed app
	Ext      Extensions
}

// Returns the privacy policy status, with default fallback
func (a *App) IsPrivacyPolicy() bool {
	if a.Privacypolicy != nil {
		return *a.Privacypolicy == 1
	}
	return false
}

// Returns the paid status, with default fallback
func (a *App) IsPaid() bool {
	if a.Paid != nil {
		return *a.Paid == 1
	}
	return false
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

// The "device" object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	Dnt            *int    // "1": Do not track
	Ua             *string // User agent
	Ip             *string // IPv4
	Geo            *Geo
	Didsha1        *string // SHA1 hashed device ID
	Didmd5         *string // MD5 hashed device ID
	Dpidsha1       *string // SHA1 hashed platform device ID
	Dpidmd5        *string // MD5 hashed platform device ID
	Ipv6           *string // IPv6
	Carrier        *string // Carrier or ISP derived from the IP address
	Language       *string // Browser language
	Make           *string // Device make
	Model          *string // Device model
	Os             *string // Device OS
	Osv            *string // Device OS version
	Js             *int    // Javascript status ("0": Disabled, "1": Enabled)
	Connectiontype *int
	Devicetype     *int
	Flashver       *string // Flash version
	Ext            map[string]string
}

// Returns the DNT status, with default fallback
func (d *Device) IsDnt() bool {
	if d.Dnt != nil {
		return *d.Dnt == 1
	}
	return false
}

// Returns the JS status, with default fallback
func (d *Device) IsJs() bool {
	if d.Js != nil {
		return *d.Js == 1
	}
	return false
}

// Returns the connection type, with default fallback
func (d *Device) ConnectionType() int {
	if d.Connectiontype != nil {
		return *d.Connectiontype
	}
	return CONN_TYPE_UNKNOWN
}

// Returns the connection type, with default fallback
func (d *Device) DeviceType() int {
	if d.Devicetype != nil {
		return *d.Devicetype
	}
	return DEVICE_TYPE_UNKNOWN
}

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
