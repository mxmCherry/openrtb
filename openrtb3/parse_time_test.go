package openrtb3_test

import (
	"time"

	. "github.com/mxmCherry/openrtb/openrtb3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseTime", func() {

	DescribeTable(
		"different date/time formats",
		func(input string, expected time.Time) {
			Expect(ParseTime(input)).To(Equal(expected))
		},

		Entry(
			"RFC3339 UTC",
			"2018-10-14T08:45:16+00:00",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.FixedZone("", 0)), // zone parsed as "nameless"
		),
		Entry(
			"RFC3339 GMT+00:12",
			"2018-10-14T08:45:16+00:12",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.FixedZone("", 12*60)), // zone parsed as "nameless"
		),
		Entry(
			"RFC3339 GMT-00:12",
			"2018-10-14T08:45:16-00:12",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.FixedZone("", -12*60)), // zone parsed as "nameless"
		),

		// https://en.wikipedia.org/wiki/ISO_8601
		Entry(
			"Date",
			"2018-10-14",
			time.Date(2018, 10, 14, 0, 0, 0, 0, time.UTC),
		),
		Entry(
			"Combined date and time in UTC 1",
			"2018-10-14T08:45:16+00:00",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.FixedZone("", 0)), // zone parsed as "nameless"
		),
		Entry(
			"Combined date and time in UTC 2",
			"2018-10-14T08:45:16Z",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.UTC),
		),
		Entry(
			"Combined date and time in UTC 3",
			"20181014T084516Z",
			time.Date(2018, 10, 14, 8, 45, 16, 0, time.UTC),
		),
	)
})

// ----------------------------------------------------------------------------
