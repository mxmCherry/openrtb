package openrtb3

import "encoding/json"

// Deal object constitutes a specific deal that was struck a priori between a seller and a buyer.
// Its presence indicates that this item is available under the terms of that deal.
type Deal struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Definition:
	//   A unique identifier for the deal.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   flr
	// Type:
	//   float
	// Definition:
	//   Minimum deal price for this item expressed in CPM.
	Flr float64 `json:"flr,omitempty"`

	// Attribute:
	//   flrcur
	// Type:
	//   string; default "USD"
	// Definition:
	//   Currency of the flr attribute specified using ISO-4217 alpha codes.
	FlrCur string `json:"flrcur,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer
	// Definition:
	//   Optional override of the overall auction type of the request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in flr is the agreed upon deal price.
	//   Additional auction types can be defined by the exchange using 500+ values.
	AT AuctionType `json:"at,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Definition:
	//   Whitelist of buyer seats allowed to bid on this deal.
	//   IDs of seats and the buyer’s customers to which they refer must be coordinated between bidders and the exchange a priori.
	//   Omission implies no restrictions.
	WSeat []string `json:"wseat,omitempty"`

	// Attribute:
	//   wadomain
	// Type:
	//   string array
	// Definition:
	//   Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal.
	//   Omission implies no restrictions.
	WADomain []string `json:"wadomain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional exchange-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
