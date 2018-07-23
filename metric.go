package openrtb

// Metric object is associated with an impression as an array of metrics. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
type Metric struct {
	Type   string    `json:"type"`
	Value  float64   `json:"value"`
	Vendor string    `json:"vendor"`
	Ext    Extension `json:"ext,omitempty"`
}
