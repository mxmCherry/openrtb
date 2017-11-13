package openrtb_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidResponse", func() {
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
			"Ad Served on Win Notice",
			"bid-response/ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"VAST XML Document Returned Inline",
			"bid-response/vast-xml-document-returned-inline.json",
			new(BidResponse)),
		Entry(
			"Direct Deal Ad Served on Win Notice",
			"bid-response/direct-deal-ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"Native Markup Returned Inline",
			"bid-response/native-markup-returned-inline.json",
			new(BidResponse)),
	)
})
