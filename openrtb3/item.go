package openrtb

import "encoding/json"

// Item object represents a unit of goods being offered for sale either on the open market or in relation to a private marketplace deal.
// The id attribute is required since there may be multiple items being offered in the same bid request and bids must reference the specific item of interest.
// This object interfaces to Layer-4 domain objects for deeper specification of the item being offered (e.g., an impression).
type Item struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Definition:
	//   A unique identifier for this item within the context of the offer (typically starts with “1” and increments).
	ID string `json:"id"`

	// Attribute:
	//   qty
	// Type:
	//   integer; default 1
	// Definition:
	//   The number of instances (i.e., “quantity”) of this item being offered (e.g., multiple identical impressions in a digital out-of-home scenario).
	Qty int `json:"qty,omitempty"`

	// Attribute:
	//   seq
	// Type:
	//   integer
	// Definition:
	//   If multiple items are offered in the same bid request, the sequence number allows for the coordinated delivery.
	Seq int `json:"seq,omitempty"`

	// Attribute:
	//   flr
	// Type:
	//   float
	// Definition:
	//   Minimum bid price for this item expressed in CPM.
	Flr float64 `json:"flr,omitempty"`

	// Attribute:
	//   flrcur
	// Type:
	//   string; default “USD”
	// Definition:
	//   Currency of the flr attribute specified using ISO-4217 alpha codes.
	FlrCur string `json:"flrcur,omitempty"`

	// Attribute:
	//   exp
	// Type:
	//   integer
	// Definition:
	//   Advisory as to the number of seconds that may elapse between auction and fulfilment.
	Exp int64 `json:"exp,omitempty"`

	// Attribute:
	//   dt
	// Type:
	//   integer
	// Definition:
	//   Timestamp when the item is expected to be fulfilled (e.g when a DOOH impression will be displayed) in Unix format (i.e., milliseconds since the epoch).
	DT int64 `json:"dt,omitempty"`

	// Attribute:
	//   dlvy
	// Type:
	//   integer; default 0
	// Definition:
	//   Item (e.g., an Ad object) delivery method required, where 0 = either method, 1 = the item must be sent as part of the transaction (e.g., by value in the bid itself, fetched by URL included in the bid), and 2 = an item previously uploaded to the exchange must be referenced by its ID.
	//   Note that if an exchange does not supported prior upload, then the default of 0 is effectively the same as 1 since there can be no items to reference.
	Dlvy int8 `json:"dlvy,omitempty"`

	// Attribute:
	//   metric
	// Type:
	//   object array
	// Definition:
	//   An array of Metric objects.
	//   Refer to Object: Metric.
	Metric []Metric `json:"metric,omitempty"`

	// Attribute:
	//   deal
	// Type:
	//   object array
	// Definition:
	//   Array of Deal objects that convey special terms applicable to this item.
	//   Refer to Object: Deal.
	Deal []Deal `json:"deal,omitempty"`

	// Attribute:
	//   private
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicator of auction eligibility to seats named in Deal objects, where 0 = all bids are accepted, 1 = bids are restricted to the deals specified and the terms thereof.
	Private int8 `json:"private,omitempty"`

	// Attribute:
	//   spec
	// Type:
	//   object; required
	// Definition:
	//   Layer-4 domain object structure that provides specifies the item being offered conforming to the specification and version referenced in openrtb.domainspec and openrtb.domainver.
	//   For AdCOM v1.x, the objects allowed here are Placement and any objects subordinate to these as specified by AdCOM.
	Spec json.RawMessage `json:"spec"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional exchange-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
