package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestResponse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Response Suite")
}
