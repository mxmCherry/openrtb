package openrtb3

import "encoding/json"

// Response object is the bid response object under the Openrtb root.
// Its id attribute is a reflection of the bid request ID.
// The bidid attribute is an optional response tracking ID for bidders.
// If specified, it will be available for use in substitution macros placed in markup and notification URLs.
// At least one Seatbid object is required, which contains at least one Bid for an item.
// Other attributes are optional.
//
// To express a “no-bid”, the most compact option is simply to return an empty response with HTTP 204.
// However, if the bidder wishes to convey a reason for not bidding, a Response object can be returned with just a reason code in the nbr attribute.
type Response struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Definition:
	//   ID of the bid request to which this is a response; must match the request.id attribute.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   bidid
	// Type:
	//   string
	// Definition:
	//   Bidder generated response ID to assist with logging/tracking.
	BidID string `json:"bidid,omitempty"`

	// Attribute:
	//   nbr
	// Type:
	//   integer
	// Definition:
	//   Reason for not bidding if applicable (see List: No-Bid Reason Codes).
	//   Note that while many exchanges prefer a simple HTTP 204 response to indicate a no-bid, responses indicating a reason code can be useful in debugging scenarios.
	NBR NoBidReason `json:"nbr,omitempty"`

	// Attribute:
	//   cur
	// Type:
	//   string; default “USD”
	// Definition:
	//   Bid currency using ISO-4217 alpha codes.
	Cur string `json:"cur,omitempty"`

	// Attribute:
	//   cdata
	// Type:
	//   string
	// Definition:
	//   Allows bidder to set data in the exchange’s cookie, which can be retrieved on bid requests (refer to cdata in Object: Request) if supported by the exchange.
	//   The string must be in base85 cookie-safe characters.
	CData string `json:"cdata,omitempty"`

	// Attribute:
	//   seatbid
	// Type:
	//   object array
	// Definition:
	//   Array of Seatbid objects; 1+ required if a bid is to be made.
	//   Refer to Object: Seatbid.
	SeatBid []SeatBid `json:"seatbid,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional demand source specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
