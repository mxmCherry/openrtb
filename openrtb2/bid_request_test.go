package openrtb2_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/mxmCherry/openrtb/v16/openrtb2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidRequest", func() {
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
			"2.5 Simple Banner",
			"bid-request/2.5/simple-banner.json",
			new(BidRequest)),
		Entry(
			"2.5 Expandable Creative",
			"bid-request/2.5/expandable-creative.json",
			new(BidRequest)),
		Entry(
			"2.5 Mobile",
			"bid-request/2.5/mobile.json",
			new(BidRequest)),
		Entry(
			"2.5 Video",
			"bid-request/2.5/video.json",
			new(BidRequest)),
		Entry(
			"2.5 PMP with Direct Deal",
			"bid-request/2.5/pmp-with-direct-deal.json",
			new(BidRequest)),
		Entry(
			"2.5 Native Ad",
			"bid-request/2.5/native-ad.json",
			new(BidRequest)),

		// 2.6
		Entry(
			"2.6 Simple Banner",
			"bid-request/2.6/simple-banner.json",
			new(BidRequest)),
		Entry(
			"2.6 Expandable Creative",
			"bid-request/2.6/expandable-creative.json",
			new(BidRequest)),
		Entry(
			"2.6 Mobile",
			"bid-request/2.6/mobile.json",
			new(BidRequest)),
		Entry(
			"2.6 Video",
			"bid-request/2.6/video.json",
			new(BidRequest)),
		Entry(
			"2.6 PMP with Direct Deal",
			"bid-request/2.6/pmp-with-direct-deal.json",
			new(BidRequest)),
	)
})
