package openrtb3

import "encoding/json"

// Bid is an OpenRTB bid object.
// A Seatbid object contains one or more Bid objects, each of which relates to a specific item in the bid request offer via the “item” attribute and constitutes an offer to buy that item for a given price.
type Bid struct {
	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Definition:
	//   Bidder generated bid ID to assist with logging/tracking.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   item
	// Type:
	//   string; required
	// Definition:
	//   ID of the item object in the related bid request; specifically item.id.
	Item string `json:"item,omitempty"`

	// Attribute:
	//   price
	// Type:
	//   float; required
	// Definition:
	//   Bid price expressed as CPM although the actual transaction is for a unit item only
	//   Note that while the type indicates float, integer math is highly recommended when handling currencies (e.g., BigDecimal in Java).
	Price float64 `json:"price,omitempty"`

	// Attribute:
	//   deal
	// Type:
	//   string
	// Definition:
	//   Reference to a deal from the bid request if this bid pertains to a private marketplace deal; specifically deal.id.
	Deal string `json:"deal,omitempty"`

	// Attribute:
	//   cid
	// Type:
	//   string
	// Definition:
	//   Campaign ID or other similar grouping of brand-related ads
	//   Typically used to increase the efficiency of audit processes.
	CID string `json:"cid,omitempty"`

	// Attribute:
	//   tactic
	// Type:
	//   string
	// Definition:
	//   Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted
	//   The specific usage and meaning of the tactic ID should be communicated between buyer and exchanges a priori.
	Tactic string `json:"tactic,omitempty"`

	// Attribute:
	//   purl
	// Type:
	//   string
	// Definition:
	//   Pending notice URL called by the exchange when a bid has been declared the winner within the scope of an OpenRTB compliant supply chain (i.e., there may still be non-compliant decisioning such as header bidding)
	// Substitution macros may be included.
	PURL string `json:"purl,omitempty"`

	// Attribute:
	//   burl
	// Type:
	//   string; recommended
	// Definition:
	//   Billing notice URL called by the exchange when a winning bid becomes billable based on exchange-specific business policy (e.g., markup rendered)
	// Substitution macros may be included.
	BURL string `json:"burl,omitempty"`

	// Attribute:
	//   lurl
	// Type:
	//   string
	// Definition:
	//   Loss notice URL called by the exchange when a bid is known to have been lost
	// Substitution macros may be included
	// Exchange-specific policy may preclude support for loss notices or the disclosure of winning clearing prices resulting in ${OPENRTB_PRICE} macros being removed (i.e., replaced with a zero-length string).
	LURL string `json:"lurl,omitempty"`

	// Attribute:
	//   exp
	// Type:
	//   integer
	// Definition:
	//   Advisory as to the number of seconds the buyer is willing to wait between auction and fulfilment.
	Exp int64 `json:"exp,omitempty"`

	// Attribute:
	//   mid
	// Type:
	//   string
	// Definition:
	//   ID to enable media to be specified by reference if previously uploaded to the exchange rather than including it by value in the domain objects.
	MID string `json:"mid,omitempty"`

	// Attribute:
	//   macro
	// Type:
	//   object array
	// Definition:
	//   Array of Macro objects that enable bid specific values to be substituted into markup; especially useful for previously uploaded media referenced via the mid attribute
	// Refer to Object: Macro.
	Macro []Macro `json:"macro,omitempty"`

	// Attribute:
	//   media
	// Type:
	//   object
	// Definition:
	//   Layer-4 domain object structure that specifies the media to be presented if the bid is won conforming to the specification and version referenced in openrtb.domainspec and openrtb.domainver
	//   For AdCOM v1.x, the objects allowed here are “Ad” and any objects subordinate thereto as specified by AdCOM.
	// Dev note:
	//   This object is implemented as ../adcom1.BidMedia type.
	Media json.RawMessage `json:"media,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional demand source specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
