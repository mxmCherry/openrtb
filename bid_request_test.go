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

		Entry(
			"Simple Banner",
			"bid-request/simple-banner.json",
			new(BidRequest)),
		Entry(
			"Expandable Creative",
			"bid-request/expandable-creative.json",
			new(BidRequest)),
		Entry(
			"Mobile",
			"bid-request/mobile.json",
			new(BidRequest)),
		Entry(
			"Video",
			"bid-request/video.json",
			new(BidRequest)),
		Entry(
			"PMP with Direct Deal",
			"bid-request/pmp-with-direct-deal.json",
			new(BidRequest)),
		Entry(
			"Native Ad",
			"bid-request/native-ad.json",
			new(BidRequest)),
	)
})
