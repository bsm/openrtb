package openrtb

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*******************************************************************
 * TEST SUITE HOOK
 *******************************************************************/

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "github.com/bsm/openrtb")
}
