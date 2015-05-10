package rtb

// Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
const (
	BidRequestTestLive = 0 // 0 = live mode
	BidRequestTestTest = 1 // 1 = test mode
)

// Auction type, where 1 = First Price, 2 = Second Price Plus.
// Exchange-specific auction types can be defined using values
// greater than 500.
const (
	BidRequestATFirstPrice      uint8 = 1 // 1 = First Price
	BidRequestATSecondPricePlus uint8 = 2 // 2 = Second Price Plus (default)
)

// Flag to indicate if Exchange can verify that the impressions
// offered represent all of the impressions available in context
// (e.g., all on the web page, all video spots such as pre/mid/post
// roll) to support road-blocking. 0 = no or unknown, 1 = yes, the
// impressions offered represent all that are available.
const (
	BidRequestAllImpsNo  uint8 = 0 // 0 = no or unknown
	BidRequestAllImpsYes uint8 = 1 // 1 = yes, the impressions offered represent all that are available.
)

// 3.2.1 Object: BidRequest
// 
// The top-level bid request object contains a globally unique bid request or auction ID. This id attribute is
// required as is at least one impression object (Section 3.2.2). Other attributes in this top-level object
// establish rules and restrictions that apply to all impressions being offered.
// 
// There are also several subordinate objects that provide detailed data to potential buyers. Among these
// are the Site and App objects, which describe the type of published media in which the impression(s)
// appear. These objects are highly recommended, but only one applies to a given bid request depending
// on whether the media is browser-based web content or a non-browser application, respectively.
type BidRequest struct {

	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   Unique ID of the bid request, provided by the exchange.
	ID string `json:"id"`

	// Attribute:
	//   imp
	// Type:
	//   object array; required
	// Description:
	//   Array of Imp objects (Section 3.2.2) representing the
	//   impressions offered. At least 1 Imp object is required.
	Imp []Imp `json:"imp"`

	// Attribute:
	//   site
	// Type:
	//   object; recommended
	// Description:
	//    Details via a Site object (Section 3.2.6) about the publisher's
	//    website. Only applicable and recommended for websites.
	Site Site `json:"site"`

	// Attribute:
	//   app
	// Type:
	//   object; recommended
	// Description:
	//    Details via an App object (Section 3.2.7) about the publisher's
	//    app (i.e. non-browser applications). Only applicable and recommended for apps.
	App App `json:"app"`

	// Attribute:
	//   device
	// Type:
	//   object; recommended
	// Description:
	//    Details via a Device object (Section 3.2.11) about the user’s device to which the impression will be delivered.
	Device Device `json:"device"`

	// Attribute:
	//   user
	// Type:
	//   object; recommended
	// Description:
	//    Details via a User object (Section 3.2.13) about the human user of the device; the advertising audience.
	User User `json:"user"`

	// Attribute:
	//   test
	// Type:
	//   integer; default 0
	// Description:
	//    Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Test uint8 `json:"test"`

	// Attribute:
	//   at
	// Type:
	//   integer; default 2
	// Description:
	//    Auction type, where 1 = First Price, 2 = Second Price Plus.
	//    Exchange-specific auction types can be defined using values
	//    greater than 500.
	AT uint8 `json:"at"`

	// Attribute:
	//   tmax
	// Type:
	//   integer
	// Description:
	//    Maximum time in milliseconds to submit a bid to avoid timeout. This value is commonly communicated offline.
	TMax int64 `json:"tmax"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   Whitelist of buyer seats allowed to bid on this deal.  Seat IDs must be
	//   communicated between bidders and the exchange a priori. Omission implies no seat restrictions.
	WSeat []string `json:"wseat"`

	// Attribute:
	//   allimps
	// Type:
	//   integer; default 0
	// Description:
	//   Flag to indicate if Exchange can verify that the impressions
	//   offered represent all of the impressions available in context
	//   (e.g., all on the web page, all video spots such as pre/mid/post
	//   roll) to support road-blocking. 0 = no or unknown, 1 = yes, the
	//   impressions offered represent all that are available.
	AllImps uint8 `json:"allimps"`

	// Attribute:
	//   cur
	// Type:
	//   string array
	// Description:
	//    Array of allowed currencies for bids on this bid request using ISO-4217 alpha codes. Recommended only if
	//    the exchange accepts multiple currencies.
	Cur []string `json:"cur"`

	// Attribute:
	//   bcat
	// Type:
	//   string array
	// Description:
	//   Blocked advertiser categories using the IAB content categories. Refer to List 5.1.
	BCat []string `json:"bcat"`

	// Attribute:
	//   badv
	// Type:
	//   string array
	// Description:
	//   Blocked advertiser categories using the IAB content categories. Refer to List 5.1.
	BAdv []string `json:"badv"`

	// Attribute:
	//   regs
	// Type:
	//   object
	// Description:
	//   Blocked advertiser categories using the IAB content categories. Refer to List 5.1. or governmental
	//   regulations in force for this request.
	Regs Regs `json:"regs"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}
