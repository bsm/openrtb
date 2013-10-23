package openrtb

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	errValidationReqId            = errors.New("openrtb parse: request ID missing")
	errValidationReqImp           = errors.New("openrtb parse: no impressions")
	errValidationReqSaA           = errors.New("openrtb parse: request has site and app")
	errValidationImpId            = errors.New("openrtb parse: impression ID missing")
	errValidationImpBoV           = errors.New("openrtb parse: impression has neither a banner nor video")
	errValidationImpBaV           = errors.New("openrtb parse: impression has banner and video")
	errValidationVideoMimes       = errors.New("openrtb parse: video has no mimes")
	errValidationVideoLinearity   = errors.New("openrtb parse: video linearity missing")
	errValidationVideoMinduration = errors.New("openrtb parse: video minduration missing")
	errValidationVideoMaxduration = errors.New("openrtb parse: video maxduration missing")
	errValidationVideoProtocol    = errors.New("openrtb parse: video protocol missing")
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

// Validates the request
func (req *Request) Valid() (bool, error) {
	if req.Id == nil {
		return false, errValidationReqId
	} else if len(req.Imp) == 0 {
		return false, errValidationReqImp
	} else if req.Site != nil && req.App != nil {
		return false, errValidationReqSaA
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
