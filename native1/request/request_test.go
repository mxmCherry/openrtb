package request_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb/v15/native1/request"

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
			"Social Context (v1.1)",
			"v1.1/social-context.json",
			new(Request)),
		Entry(
			"Content Context (v1.1)",
			"v1.1/content-context.json",
			new(Request)),

		Entry(
			"Social Context (v1.2)",
			"v1.2/social-context.json",
			new(Request)),
		Entry(
			"Content Context (v1.2)",
			"v1.2/content-context.json",
			new(Request)),
		Entry(
			"Third-party (v1.2)",
			"v1.2/third-party.json",
			new(Request)),
	)
})
