package openrtb3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOpenrtb3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Openrtb3 Suite")
}
