package rtb

// 3.2.2 Object: Imp
//
// This object describes an ad placement or impression being auctioned. A single bid request can include
// multiple Imp objects, a use case for which might be an exchange that supports selling all ad positions on
// a given page. Each Imp object has a required ID so that bids can reference them individually.
//
// The presence of Banner (Section 3.2.3), Video (Section 3.2.4), and/or Native (Section 3.2.5) objects
// subordinate to the Imp object indicates the type of impression being offered. The publisher can choose
// one such type which is the typical case or mix them at their discretion. However, any given bid for the
// impression must conform to one of the offered types.
type Imp struct {

	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   A unique identifier for this impression within the context of the bid request (typically, starts
	//   with 1 and increments.
	ID string `json:"id"`

	// Attribute:
	//   banner
	// Type:
	//   object
	// Description:
	//   A Banner object (Section 3.2.3); required if this impression is
	//   offered as a banner ad opportunity.
	Banner *Banner `json:"banner"`

	// Attribute:
	//   video
	// Type:
	//   object
	// Description:
	//   A Video object (Section 3.2.4); required if this impression is
	//   offered as a video ad opportunity.
	Video *Video `json:"video"`

	// Attribute:
	//   native
	// Type:
	//   object
	// Description:
	//   A Native object (Section 3.2.5); required if this impression is
	//   offered as a native ad opportunity
	Native *Native `json:"native"`

	// Attribute:
	//   displaymanager
	// Type:
	//   string
	// Description:
	//   Name of ad mediation partner, SDK technology, or player
	//   responsible for rendering ad (typically video or mobile). Used
	//   by some ad servers to customize ad code by partner.
	//   Recommended for video and/or apps.
	DisplayManager string `json:"displaymanager"`

	// Attribute:
	//   instl
	// Type:
	//   int; default 0
	// Description:
	//   1 = the ad is interstitial or full screen, 0 = not interstitial.
	Instl uint8 `json:"instl"`

	// Attribute:
	//   tagid
	// Type:
	//   string
	// Description:
	//   Identifier for specific ad placement or ad tag that was used to
	//   initiate the auction. This can be useful for debugging of any
	//   issues, or for optimization by the buyer.
	TagID string `json:"tagid"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for this impression expressed in CPM.
	BidFloor float64 `json:"bidfloor"`

	// Attribute:
	//   bidfloorcur
	// Type:
	//   string; default “USD”
	// Description:
	//   Currency specified using ISO-4217 alpha codes. This may be
	//   different from bid currency returned by bidder if this is
	//   allowed by the exchange.
	BidFloorCur float64 `json:"bidfloorcur"`

	// Attribute:
	//   secure
	// Type:
	//   integer
	// Description:
	//   Flag to indicate if the impression requires secure HTTPS URL
	//   creative assets and markup, where 0 = non-secure, 1 = secure.
	//   If omitted, the secure state is unknown, but non-secure HTTP
	//   support can be assumed.
	Secure int8 `json:"secure"`

	// Attribute:
	//   iframebuster
	// Type:
	//   string array
	// Description:
	//   Array of exchange-specific names of supported iframe busters.
	IframeBuster []string `json:"iframebuster"`

	// Attribute:
	//   pmp
	// Type:
	//   object
	// Description:
	//   A Pmp object (Section 3.2.17) containing any private marketplace deals in effect for this impression.
	PMP *PMP `json:"pmp"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}
