package openrtb

import (
	"encoding/json"
	"strconv"
)

// NumberOrString attempts to fix OpenRTB incompatibilities
// of exchanges. On decoding, it can handle numbers and strings.
// On encoding, it will generate a number, as intended by the
// standard.
type NumberOrString int

// UnmarshalJSON implements json.Unmarshaler
func (n *NumberOrString) UnmarshalJSON(data []byte) (err error) {
	var v int

	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		err = json.Unmarshal(data[1:len(data)-1], &v)
	} else {
		err = json.Unmarshal(data, &v)
	}
	if err != nil {
		return err
	}

	*n = NumberOrString(v)
	return nil
}

// StringOrNumber attempts to fix OpenRTB incompatibilities
// of exchanges. On decoding, it can handle numbers and strings.
// On encoding, it will generate a string, as intended by the
// standard.
type StringOrNumber string

// UnmarshalJSON implements json.Unmarshaler
func (n *StringOrNumber) UnmarshalJSON(data []byte) error {
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*n = StringOrNumber(v)
	} else {
		var v int
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*n = StringOrNumber(strconv.Itoa(v))
	}
	return nil
}
