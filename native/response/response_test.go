package response_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb/native/response"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
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
			"Clickout (v1.1)",
			"v1.1/clickout.json",
			new(Response)),
		Entry(
			"Video (v1.1)",
			"v1.1/video.json",
			new(Response)),

		Entry(
			"Clickout (v1.2)",
			"v1.2/clickout.json",
			new(Response)),
		Entry(
			"Video (v1.2)",
			"v1.2/video.json",
			new(Response)),
		Entry(
			"Third-party (v1.2)",
			"v1.2/third-party.json",
			new(Response)),
	)
})
