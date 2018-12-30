package openrtb2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpenrtb2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Openrtb2 Suite")
}
