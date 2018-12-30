package openrtb3

import "encoding/json"

// SeatBid is an OpenRTB seatbid object.
// A bid response can contain multiple Seatbid objects, each on behalf of a different buyer seat and each containing one or more individual bids.
// If multiple items are presented in the request offer, the package attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is interested in winning any only if it can win them all as a group.
type SeatBid struct {
	// Attribute:
	//   seat
	// Type:
	//   string, recommended
	// Definition:
	//   ID of the buyer seat on whose behalf this bid is made.
	Seat string `json:"seat,omitempty"`

	// Attribute:
	//   package
	// Type:
	//   integer; default 0
	// Definition:
	//   For offers with multiple items, this flag Indicates if the bidder is willing to accept wins on a subset of bids or requires the full group as a package, where 0 = individual wins accepted; 1 = package win or loss only.
	Package int8 `json:"package,omitempty"`

	// Attribute:
	//   bid
	// Type:
	//   object array; required
	// Definition:
	//   Array of 1+ Bid objects each related to an item.
	//   Multiple bids can relate to the same item.
	//   Refer to Object: Bid.
	Bid []Bid `json:"bid,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional demand source specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
