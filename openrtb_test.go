package openrtb_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/mxmCherry/openrtb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestOpenRTB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenRTB")
}

// ----------------------------------------------------------------------------

var _ = DescribeTable(
	"Marshalling",
	func(filename string, subject interface{}) {

		expected, err := ioutil.ReadFile(filepath.Join("testdata", filename))
		Expect(err).NotTo(HaveOccurred())

		Expect(json.Unmarshal(expected, subject)).To(Succeed())

		actual, err := json.Marshal(subject)
		Expect(err).NotTo(HaveOccurred())

		Expect(actual).To(MatchJSON(expected))
	},
	Entry(
		"Bid Request - Simple Banner",
		"bid-request-simple-banner.json",
		new(BidRequest)),
	Entry(
		"Bid Request - Expandable Creative",
		"bid-request-expandable-creative.json",
		new(BidRequest)),
	Entry(
		"Bid Request - Mobile",
		"bid-request-mobile.json",
		new(BidRequest)),
	Entry(
		"Bid Request - Video",
		"bid-request-video.json",
		new(BidRequest)),
	Entry(
		"Bid Request - PMP with Direct Deal",
		"bid-request-pmp-with-direct-deal.json",
		new(BidRequest)),
	Entry(
		"Bid Request - Native Ad",
		"bid-request-native-ad.json",
		new(BidRequest)),

	Entry(
		"Bid Response - Ad Served on Win Notice",
		"bid-response-ad-served-on-win-notice.json",
		new(BidResponse)),
	Entry(
		"Bid Response - VAST XML Document Returned Inline",
		"bid-response-vast-xml-document-returned-inline.json",
		new(BidResponse)),
	Entry(
		"Bid Response - Direct Deal Ad Served on Win Notice",
		"bid-response-direct-deal-ad-served-on-win-notice.json",
		new(BidResponse)),
	Entry(
		"Bid Response - Native Markup Returned Inline",
		"bid-response-native-markup-returned-inline.json",
		new(BidResponse)),
)
