package adcom1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAdcom1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adcom1 Suite")
}
