package openrtb

// 4.2.3 Object: Bid
//
// A SeatBid object contains one or more Bid objects, each of which relates to a specific impression in the
// bid request via the impid attribute and constitutes an offer to buy that impression for a given price.
type Bid struct {

	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   Bidder generated bid ID to assist with logging/tracking.
	ID string `json:"id"`

	// Attribute:
	//   impid
	// Type:
	//   string; required
	// Description:
	//   ID of the Imp object in the related bid request.
	ImpID string `json:"impid"`

	// Attribute:
	//   price
	// Type:
	//   float; required
	// Description:
	//   Bid price expressed as CPM although the actual transaction is
	//   for a unit impression only. Note that while the type indicates
	//   float, integer math is highly recommended when handling
	//   currencies (e.g., BigDecimal in Java).
	Price float64 `json:"price"`

	// Attribute:
	//   adid
	// Type:
	//   string
	// Description:
	//   ID of a preloaded ad to be served if the bid wins.
	AdID string `json:"adid,omitempty"`

	// Attribute:
	//   nurl
	// Type:
	//   string
	// Description:
	//   Win notice URL called by the exchange if the bid wins; optional
	//   means of serving ad markup.
	NURL string `json:"nurl,omitempty"`

	// Attribute:
	//   adm
	// Type:
	//   string
	// Description:
	//   Optional means of conveying ad markup in case the bid wins;
	//   supersedes the win notice if markup is included in both.
	AdM string `json:"adm,omitempty"`

	// Attribute:
	//   adomain
	// Type:
	//   string array
	// Description:
	//   Advertiser domain for block list checking (e.g., “ford.com”).
	//   This can be an array of for the case of rotating creatives.
	//   Exchanges can mandate that only one domain is allowed.
	ADomain []string `json:"adomain,omitempty"`

	// Attribute:
	//   bundle
	// Type:
	//   string
	// Description:
	//   Bundle or package name (e.g., com.foo.mygame) of the app
	//   being advertised, if applicable; intended to be a unique ID
	//   across exchanges.
	Bundle string `json:"bundle,omitempty"`

	// Attribute:
	//   iurl
	// Type:
	//   string
	// Description:
	//   URL without cache-busting to an image that is representative
	//   of the content of the campaign for ad quality/safety checking.
	IURL string `json:"iurl,omitempty"`

	// Attribute:
	//   cid
	// Type:
	//   string
	// Description:
	//   Campaign ID to assist with ad quality checking; the collection
	//   of creatives for which iurl should be representative.
	CID string `json:"cid,omitempty"`

	// Attribute:
	//   crid
	// Type:
	//   string
	// Description:
	//   Creative ID to assist with ad quality checking.
	CrID string `json:"crid,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   IAB content categories of the creative. Refer to List 5.1.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   attr
	// Type:
	//   integer array
	// Description:
	//   Set of attributes describing the creative. Refer to List 5.3.
	Attr []int8 `json:"attr,omitempty"`

	// Attribute:
	//   dealid
	// Type:
	//   string
	// Description:
	//   Reference to the deal.id from the bid request if this bid
	//   pertains to a private marketplace direct deal.
	DealID string `json:"dealid,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Height of the creative in pixels.
	H uint64 `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Width of the creative in pixels.
	W uint64 `json:"w,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for bidder-specific extensions to OpenRTB
	Ext Ext `json:"ext,omitempty"`
}
