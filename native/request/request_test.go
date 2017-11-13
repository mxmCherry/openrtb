package request_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb/native/request"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	DescribeTable(
		"Marshaling",

		func(filename string, subject interface{}) {
			expected, err := ioutil.ReadFile(filepath.Join("testdata", filename))
			Expect(err).NotTo(HaveOccurred())

			Expect(json.Unmarshal(expected, subject)).To(Succeed())

			actual, err := json.Marshal(subject)
			Expect(err).NotTo(HaveOccurred())

			Expect(actual).To(MatchJSON(expected))
		},

		Entry(
			"Social Context",
			"social-context.json",
			new(Request)),
		Entry(
			"Content Context",
			"content-context.json",
			new(Request)),
	)
})
