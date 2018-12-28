package openrtb3

import "encoding/json"

// Request object contains a globally unique bid request ID.
// This id attribute is required as is an Item array with at least one object (i.e., at least one item for sale).
// Other attributes establish rules and restrictions that apply to all items being offered.
// This object also interfaces to Layer-4 domain objects for context such as the user, device, site or app, etc.
type Request struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Definition:
	//   Unique ID of the bid request; provided by the exchange.
	ID string `json:"id"`

	// Attribute:
	//   test
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Test int8 `json:"test,omitempty"`

	// Attribute:
	//   tmax
	// Type:
	//   integer
	// Definition:
	//   Maximum time in milliseconds the exchange allows for bids to be received including Internet latency to avoid timeout.
	//   This value supersedes any a priori guidance from the exchange.
	//   If an exchange acts as an intermediary, it should decrease the outbound tmax value from what it received to account for its latency and the additional internet hop.
	TMax int64 `json:"tmax,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer; default 2
	// Definition:
	//   Auction type, where 1 = First Price, 2 = Second Price Plus.
	//   Values greater than 500 can be used for exchange-specific auction types.
	AT AuctionType `json:"at,omitempty"`

	// Attribute:
	//   cur
	// Type:
	//   string array; default [“USD”]
	// Definition:
	//   Array of accepted currencies for bids on this bid request using ISO-4217 alpha codes.
	//   Recommended if the exchange accepts multiple currencies.
	//   If omitted, the single currency of “USD” is assumed.
	Cur []string `json:"cur,omitempty"`

	// Attribute:
	//   seat
	// Type:
	//   string array
	// Definition:
	//   Restriction list of buyer seats for bidding on this item.
	//   Knowledge of buyer’s customers and their seat IDs must be coordinated between parties a priori.
	//   Omission implies no restrictions.
	Seat []string `json:"seat,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   integer; default 1
	// Definition:
	//   Flag that determines the restriction interpretation of the seat array, where 0 = block list, 1 = whitelist.
	WSeat int8 `json:"wseat,omitempty"`

	// Attribute:
	//   cdata
	// Type:
	//   string
	// Definition:
	//   Allows bidder to retrieve data set on its behalf in the exchange’s cookie (refer to cdata in Object: Response) if supported by the exchange.
	//   The string must be in base85 cookie-safe characters.
	CData string `json:"cdata,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   object
	// Definition:
	//   A Source object that provides data about the inventory source and which entity makes the final decision
	//   Refer to Object: Source.
	Source *Source `json:"source,omitempty"`

	// Attribute:
	//   item
	// Type:
	//   object array; required
	// Definition:
	//   Array of Item objects (at least one) that constitute the set of goods being offered for sale.
	//   Refer to Object: Item.
	Item []Item `json:"item"`

	// Attribute:
	//   package
	// Type:
	//   integer
	// Definition:
	//   Flag to indicate if the Exchange can verify that the items offered represent all of the items available in context (e.g., all impressions on a web page, all video spots such as pre/mid/post roll) to support road-blocking, where 0 = no, 1 = yes.
	Package int8 `json:"package,omitempty"`

	// Attribute:
	//   context
	// Type:
	//   object; recommended
	// Definition:
	//   Layer-4 domain object structure that provides context for the items being offered conforming to the specification and version referenced in openrtb.domainspec and openrtb.domainver.
	//   For AdCOM v1.x, the objects allowed here all of which are optional are one of the DistributionChannel subtypes (i.e., Site, App, or Dooh), User, Device, Regs, Restrictions, and any objects subordinate to these as specified by AdCOM.
	Context json.RawMessage `json:"context,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional exchange-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
