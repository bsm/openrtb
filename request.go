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

// Parses an OpenRTB Request struct from an io.Reader
// Optionally validates and applies defaults to the Request object (recommended)
func ParseRequest(reader io.Reader, validate bool) (req *Request, err error) {
	dec := json.NewDecoder(reader)
	if err = dec.Decode(&req); err != nil {
		return nil, err
	}
	return processReq(req, validate)
}

// Parses an OpenRTB Request from bytes
// Optionally validates and applies defaults to the Request object (recommended)
func ParseRequestBytes(data []byte, validate bool) (req *Request, err error) {
	if err = json.Unmarshal(data, &req); err != nil {
		return nil, err
	}
	return processReq(req, validate)
}

// ------------------- Request Helpers ----------------------

func processReq(req *Request, validate bool) (*Request, error) {
	if validate {
		if err := validateReq(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

func validateReq(req *Request) error {

	// Validations
	if req.Id == nil {
		return errValidationReqId
	} else if len(req.Imp) == 0 {
		return errValidationReqImp
	} else if req.Site != nil && req.App != nil {
		return errValidationReqSaA
	} else if req.Site != nil {
		if err := validateSite(req.Site); err != nil {
			return err
		}
	} else if req.App != nil {
		if err := validateApp(req.App); err != nil {
			return err
		}
	} else if req.Device != nil {
		if err := validateDevice(req.Device); err != nil {
			return err
		}
	}
	for _, imp := range req.Imp {
		if err := validateImp(&imp); err != nil {
			return err
		}
	}

	// Defaults
	if req.At == nil {
		req.At = new(int)
		*req.At = 2
	}

	return nil
}

func validateImp(imp *Impression) error {

	// Validations
	if imp.Id == nil {
		return errValidationImpId
	} else if imp.Banner != nil && imp.Video != nil {
		return errValidationImpBaV
	} else if imp.Banner != nil {
		if err := validateBanner(imp.Banner); err != nil {
			return err
		}
	} else if imp.Video != nil {
		if err := validateVideo(imp.Video); err != nil {
			return err
		}
	} else {
		return errValidationImpBoV
	}

	// Defaults
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

	return nil
}

func validateBanner(banner *Banner) error {

	// Defaults
	if banner.Topframe == nil {
		banner.Topframe = new(int)
		*banner.Topframe = 0
	}
	if banner.Pos == nil {
		banner.Pos = new(int)
		*banner.Pos = AD_POS_UNKNOWN
	}

	return nil
}

func validateVideo(video *Video) error {

	// Validations
	if len(video.Mimes) == 0 {
		return errValidationVideoMimes
	} else if video.Linearity == nil {
		return errValidationVideoLinearity
	} else if video.Minduration == nil {
		return errValidationVideoMinduration
	} else if video.Maxduration == nil {
		return errValidationVideoMaxduration
	} else if video.Protocol == nil {
		return errValidationVideoProtocol
	}

	// Defaults
	if video.Sequence == nil {
		video.Sequence = new(int)
		*video.Sequence = 1
	}
	if video.Boxingallowed == nil {
		video.Boxingallowed = new(int)
		*video.Boxingallowed = 1
	}
	if video.Pos == nil {
		video.Pos = new(int)
		*video.Pos = AD_POS_UNKNOWN
	}

	return nil
}

func validateSite(site *Site) error {

	// Defaults
	if site.Privacypolicy == nil {
		site.Privacypolicy = new(int)
		*site.Privacypolicy = 0
	}

	return nil
}

func validateApp(app *App) error {

	// Defaults
	if app.Privacypolicy == nil {
		app.Privacypolicy = new(int)
		*app.Privacypolicy = 0
	}
	if app.Paid == nil {
		app.Paid = new(int)
		*app.Paid = 0
	}

	return nil
}

func validateDevice(device *Device) error {

	// Defaults
	if device.Dnt == nil {
		device.Dnt = new(int)
		*device.Dnt = 0
	}
	if device.Js == nil {
		device.Js = new(int)
		*device.Js = 0
	}
	if device.Connectiontype == nil {
		device.Connectiontype = new(int)
		*device.Connectiontype = CONN_TYPE_UNKNOWN
	}
	if device.Devicetype == nil {
		device.Devicetype = new(int)
		*device.Devicetype = DEVICE_TYPE_UNKNOWN
	}

	return nil
}
