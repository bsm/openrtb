package request_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/octoclick/openrtb-easyjson"
	. "github.com/octoclick/openrtb-easyjson/native/request"
)

func TestRequest(t *testing.T) {
	var subject *Request
	if err := fixture("testdata/request1.json", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Request{
		Version:          "1.1",
		ContextTypeID:    ContextTypeSocial,
		ContextSubTypeID: ContextSubTypeSocial,
		PlacementTypeID:  PlacementTypeID(11),
		PlacementCount:   1,
		Sequence:         2,
		Assets: []Asset{
			{ID: 123, Required: 1, Title: &Title{Length: 140}},
			{ID: 128, Image: &Image{TypeID: ImageTypeMain, WidthMin: 836, HeightMin: 627, Width: 1000, Height: 800, MIMEs: []string{"image/jpg"}}},
			{ID: 126, Required: 1, Data: &Data{TypeID: DataTypeSponsored, Length: 25}},
			{ID: 127, Required: 1, Data: &Data{TypeID: DataTypeDesc, Length: 140}},
			{ID: 4, Video: &Video{MinDuration: 15, MaxDuration: 30, Protocols: []openrtb.Protocol{openrtb.ProtocolVAST2, openrtb.ProtocolVAST3}, MIMEs: []string{"video/mp4"}}},
		},
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}

func fixture(path string, v interface{}) error {
	bin, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(bin, v)
}
