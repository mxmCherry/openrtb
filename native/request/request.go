// Package request provides OpenRTB Dynamic Native Ads API Specification Version 1.1
// section 4 Native Ad Request Markup Details types:
// https://www.iab.com/guidelines/real-time-bidding-rtb-project/
// https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-Native-Ads-Specification-1-1_2016.pdf
package request

// 4.1 Native Markup Request Object
//
// The Native Object defines the native advertising opportunity available for bid via this bid request.
// It will be included as a JSON-encoded string in the bid request’s imp.native field or as a direct JSON object, depending on the choice of the exchange.
// While OpenRTB 2.3/2.4 supports only JSON-encoded strings, many exchanges have implemented a formal object.
// Check with your integration docs.
//
// The Default column dictates how optional parameters should be interpreted if explicit values are not provided.
type Request struct {

	// Field:
	//   ver
	// Scope:
	//   optional
	// Type:
	//   string
	// Default:
	//   1.1
	// Description:
	//   Version of the Native Markup version in use.
	Ver string `json:"ver,omitempty"`

	// Field:
	//   layout
	// Scope:
	//   recommended in 1.0, to be deprecated
	// Type:
	//   integer
	// Description:
	//   The Layout ID of the native ad unit.
	//   See the Table of Layout IDs below.
	Layout Layout `json:"layout,omitempty"`

	// Field:
	//   adunit
	// Scope:
	//   recommended in 1.0, to be deprecated
	// Type:
	//   integer
	// Description:
	//   The Ad unit ID of the native ad unit.
	//   See Table of Ad Unit IDs below for a list of supported core ad units.
	AdUnit AdUnit `json:"adunit,omitempty"`

	// Field:
	//   context
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Description:
	//   The context in which the ad appears.
	//   See Table of Context IDs below for a list of supported context types.
	Context ContextType `json:"context,omitempty"`

	// Field:
	//   contextsubtype
	// Scope:
	//   optional
	// Type:
	//   integer
	// Description:
	//   A more detailed context in which the ad appears.
	//   See Table of Context SubType IDs below for a list of supported context subtypes.
	ContextSubType ContextSubType `json:"contextsubtype,omitempty"`

	// Field:
	//   plcmttype
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Description:
	//   The design/format/layout of the ad unit being offered.
	//   See Table of Placement Type IDs below for a list of supported placement types.
	PlcmtType PlacementType `json:"plcmttype,omitempty"`

	// Field:
	//   plcmtcnt
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   1
	// Description:
	//   The number of identical placements in this Layout.
	//   Refer Section 8.1 Multiplacement Bid Requests for further detail.
	PlcmtCnt int64 `json:"plcmtcnt,omitempty"`

	// Field:
	//   seq
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   0
	// Description:
	//   0 for the first ad, 1 for the second ad, and so on.
	//   Note this would generally NOT be used in combination with plcmtcnt - either you are auctioning multiple identical placements (in which case plcmtcnt>1, seq=0) or you are holding separate auctions for distinct items in the feed (in which case plcmtcnt=1, seq=>=1)
	Seq int64 `json:"seq,omitempty"`

	// Field:
	//   assets
	// Scope:
	//   required
	// Type:
	//   array of objects
	// Description:
	//   An array of Asset Objects.
	//   Any bid response must comply with the array of elements expressed in the bid request.
	Assets []Asset `json:"assets"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	// This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext RawJSON `json:"ext,omitempty"`
}
