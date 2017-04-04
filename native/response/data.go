package response

import "github.com/Upliner/openrtb"

type Data struct {
	Label string            `json:"label,omitempty"` // The optional formatted string name of the data type to be displayed
	Value string            `json:"value"`           // The formatted string of data to be displayed. Can contain a formatted value such as “5 stars” or “$10” or “3.4 stars out of 5”
	Ext   openrtb.Extension `json:"ext,omitempty"`
}
