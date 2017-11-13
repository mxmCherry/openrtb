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
			"Clickout",
			"clickout.json",
			new(Response)),
		Entry(
			"Video",
			"video.json",
			new(Response)),
	)
})
