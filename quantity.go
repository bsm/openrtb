package openrtb

import (
	"encoding/json"
	"errors"
)

var (
	ErrMissingMultiplier        = errors.New("openrtb: qty.multiplier is required")
	ErrMissingMeasurementVendor = errors.New("openrtb: qty.vendor is required when qty.sourcetype is 1 (Measurement Vendor)")
)

type MeasurementSourceType int

const (
	MeasurementSourceTypeUnknown           MeasurementSourceType = 0
	MeasurementSourceTypeMeasurementVendor MeasurementSourceType = 1
	MeasurementSourceTypePublisher         MeasurementSourceType = 2
	MeasurementSourceTypeExchange          MeasurementSourceType = 3
)

type Quantity struct {
	Multiplier float64               `json:"multiplier"`
	SourceType MeasurementSourceType `json:"sourcetype,omitempty"`
	Vendor     string                `json:"vendor,omitempty"`
	Ext        *json.RawMessage       `json:"ext,omitempty"`
}

func (qty *Quantity) Validate() error {
	if qty.Multiplier == 0 {
		return ErrMissingMultiplier
	}
	if qty.SourceType == MeasurementSourceTypeMeasurementVendor && qty.Vendor == "" {
		return ErrMissingMeasurementVendor
	}
	return nil
}
