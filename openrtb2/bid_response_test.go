package openrtb2_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb/v17/openrtb2"

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

		// 2.5
		Entry(
			"2.5 Ad Served on Win Notice",
			"bid-response/2.5/ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"2.5 VAST XML Document Returned Inline",
			"bid-response/2.5/vast-xml-document-returned-inline.json",
			new(BidResponse)),
		Entry(
			"2.5 Direct Deal Ad Served on Win Notice",
			"bid-response/2.5/direct-deal-ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"2.5 Native Markup Returned Inline",
			"bid-response/2.5/native-markup-returned-inline.json",
			new(BidResponse)),

		// 2.6
		Entry(
			"2.6 Ad Served on Win Notice",
			"bid-response/2.6/ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"2.6 VAST XML Document Returned Inline",
			"bid-response/2.6/vast-xml-document-returned-inline.json",
			new(BidResponse)),
		Entry(
			"2.6 Direct Deal Ad Served on Win Notice",
			"bid-response/2.6/direct-deal-ad-served-on-win-notice.json",
			new(BidResponse)),
		Entry(
			"2.6 Native Markup Returned Inline",
			"bid-response/2.6/native-markup-returned-inline.json",
			new(BidResponse)),
	)
})
