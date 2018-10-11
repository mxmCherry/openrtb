package openrtb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpenrtb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Openrtb Suite")
}
