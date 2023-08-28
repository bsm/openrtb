package openrtb_test

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

func fixture(fname string, v interface{}) error {
	f, err := os.Open(filepath.Join("testdata", fname+".json"))
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}
