package openrtb_test

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/UnityTech/openrtb"
)

func TestExtension(t *testing.T) {
	var _ json.Marshaler = (Extension)(nil)
	var _ json.Unmarshaler = (*Extension)(nil)

	t.Run("should marshal JSON", func(t *testing.T) {
		subject := Extension(`{"foo":"bar"}`)

		got, err := json.Marshal(subject)
		if err != nil {
			t.Fatal(err.Error())
		}
		exp := []byte(`{"foo":"bar"}`)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v, got %+v", exp, got)
		}
	})

	t.Run("should unmarshal JSON", func(t *testing.T) {
		subject := Extension([]byte{})
		err := json.Unmarshal([]byte(`{"foo":"bar"}`), &subject)
		if err != nil {
			t.Fatal(err.Error())
		}
		exp := Extension(`{"foo":"bar"}`)
		if got := subject; !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v, got %+v", exp, got)
		}
	})
}
