package openrtb

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openrtb")
}

func iptr(n int) *int       { return &n }
func sptr(s string) *string { return &s }

func fixture(fname string, v interface{}) error {
	data, err := ioutil.ReadFile(filepath.Join("testdata", fname+".json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
