package openrtb_test

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func fixture(fname string, v interface{}) error {
	f, err := os.Open(filepath.Join("testdata", fname+".json"))
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}
