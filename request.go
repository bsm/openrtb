package openrtb

import (
	"encoding/json"
	"errors"
	"io"
)

// The top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type Request struct {
	Id      *string      `json:"id"` // Unique ID of the bid request
	Imp     []Impression `json:"imp,omitempty"`
	Site    *Site        `json:"site,omitempty"`
	App     *App         `json:"app,omitempty"`
	Device  *Device      `json:"device,omitempty"`
	User    *User        `json:"user,omitempty"`
	At      *int         `json:"at"`                // Auction type, Default: 2 ("1": first price auction, "2": then second price auction)
	Tmax    *int         `json:"tmax,omitempty"`    // Maximum amount of time in milliseconds to submit a bid
	Wseat   []string     `json:"wseat,omitempty"`   // Array of buyer seats allowed to bid on this auction
	Allimps *int         `json:"allimps,omitempty"` // Flag to indicate whether exchange can verify that all impressions offered represent all of the impressions available in context, Default: 0
	Cur     []string     `json:"cur,omitempty"`     // Array of allowed currencies
	Bcat    []string     `json:"bcat,omitempty"`    // Blocked Advertiser Categories.
	Badv    []string     `json:"badv,omitempty"`    // Array of strings of blocked toplevel domains of advertisers
	Regs    *Regulations `json:"regs,omitempty"`
	Ext     Extensions   `json:"ext,omitempty"`

	Pmp *Pmp `json:"pmp,omitempty"` // DEPRECATED: kept for backwards compatibility
}

// Parses an OpenRTB Request struct from an io.Reader
// Optionally validates and applies defaults to the Request object (recommended)
func ParseRequest(reader io.Reader) (req *Request, err error) {
	dec := json.NewDecoder(reader)
	if err = dec.Decode(&req); err != nil {
		return nil, err
	}
	return req.WithDefaults(), nil
}

// Parses an OpenRTB Request from bytes
// Optionally validates and applies defaults to the Request object (recommended)
func ParseRequestBytes(data []byte) (req *Request, err error) {
	if err = json.Unmarshal(data, &req); err != nil {
		return nil, err
	}
	return req.WithDefaults(), nil
}

// Validation errors
var (
	ErrInvalidReqID  = errors.New("openrtb parse: request ID missing")
	ErrInvalidReqImp = errors.New("openrtb parse: no impressions")
	ErrInvalidReqSrc = errors.New("openrtb parse: request has site and app")
)

// Validates the request
func (req *Request) Valid() (bool, error) {
	if req.Id == nil {
		return false, ErrInvalidReqID
	} else if len(req.Imp) == 0 {
		return false, ErrInvalidReqImp
	} else if req.Site != nil && req.App != nil {
		return false, ErrInvalidReqSrc
	}

	for _, imp := range req.Imp {
		if ok, err := (&imp).Valid(); !ok {
			return ok, err
		}
	}

	return true, nil
}

// Applies defaults
func (req *Request) WithDefaults() *Request {
	if req.At == nil {
		req.At = new(int)
		*req.At = 2
	}

	if req.Site != nil {
		req.Site.WithDefaults()
	}
	if req.App != nil {
		req.App.WithDefaults()
	}
	if req.Device != nil {
		req.Device.WithDefaults()
	}
	for i, imp := range req.Imp {
		req.Imp[i] = *(&imp).WithDefaults()
	}

	return req
}

// Set the ID
func (req *Request) SetId(id string) *Request {
	if req.Id == nil {
		req.Id = new(string)
	}
	*req.Id = id
	return req
}

// Set the Site
func (req *Request) SetSite(site *Site) *Request {
	req.Site = site
	return req
}

// Set the App
func (req *Request) SetApp(app *App) *Request {
	req.App = app
	return req
}
