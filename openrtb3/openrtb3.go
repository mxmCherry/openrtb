// Package openrtb3 holds OpenRTB 3.0 types implementation
// (https://github.com/InteractiveAdvertisingBureau/openrtb)
package openrtb3

import "encoding/json"

// Body is a top-level wrapper for OpenRTB requests.
type Body struct {
	OpenRTB OpenRTB `json:"openrtb"`
}

// ----------------------------------------------------------------------------

// OpenRTB is the top-level object is the root for both request and response payloads.
// It includes versioning information and references to the Layer-4 domain model on which transactions are based.
// By default, the domain model used by OpenRTB is the Advertising Common Object Model (AdCOM).
type OpenRTB struct {
	Ver        string    `json:"ver,omitempty"`        // Version of the Layer-3 OpenRTB specification (e.g., "3.0").
	DomainSpec string    `json:"domainspec,omitempty"` // Default: "adcom" (const DefaultOpenRTBDomainSpec). Identifier of the Layer-4 domain model used to define items for sale, media associated with bids, etc.
	DomainVer  string    `json:"domainver"`            // Required. Specification version of the Layer-4 domain model referenced in the “domainspec” attribute.
	Request    *Request  `json:"request,omitempty"`    // Required only for request payloads. Bid request container.
	Response   *Response `json:"response,omitempty"`   // Required only for response payloads. Bid response container.

	// Some of these attributes are optional.
	// The "ver" attribute, for example, indicates the OpenRTB specification version to which this payload conforms.
	// This is also conveyed in Layer-1 via an HTTP header.
	// Its utility here is more to assist in diagnostics by making the payload more self-documenting outside the context of a runtime transaction.
	//
	// The "domainver" attribute, however, does have runtime utility since the structures of Layer-4 objects may vary over time based on their specification versions.
	// This attribute can assist in invoking the correct domain object parser or unmarshalling code.
}

// ----------------------------------------------------------------------------

// Request object contains a globally unique bid request ID.
// This "id" attribute is required as is an Item array with at least one object (i.e., at least one item for sale).
// Other attributes establish rules and restrictions that apply to all items being offered.
// This object also interfaces to Layer-4 domain objects for context such as the user, device, site or app, etc.
type Request struct {
	ID      string          `json:"id"`                // Required. Unique ID of the bid request; provided by the exchange.
	Test    Bool            `json:"test,omitempty"`    // Default: 0 (false). Indicator of test mode in which auctions are not billable, where 0 (false) = live mode, 1 (true) = test mode.
	TMax    int64           `json:"tmax,omitempty"`    // Maximum time in milliseconds the exchange allows for bids to be received including Internet latency to avoid timeout. This value supersedes any a priori guidance from the exchange. If an exchange acts as an intermediary, it should decrease the outbound "tmax" value from what it received to account for its latency and the additional internet hop.
	AT      int8            `json:"at,omitempty"`      // Default: 2. Auction type, where 1 = First Price, 2 = Second Price Plus. Values greater than 500 can be used for exchange-specific auction types.
	Cur     []string        `json:"cur,omitempty"`     // Default: ["USD"]. Array of accepted currencies for bids on this bid request using ISO-4217 alpha codes. Recommended if the exchange accepts multiple currencies. If omitted, the single currency of "USD" is assumed.
	Seat    []string        `json:"seat,omitempty"`    // Restriction list of buyer seats for bidding on this item. Knowledge of buyer’s customers and their seat IDs must be coordinated between parties a priori. Omission implies no restrictions.
	WSeat   *Bool           `json:"wseat,omitempty"`   // Default: 1 (true). Flag that determines the restriction interpretation of the “seat” array, where 0 (false) = block list, 1 (true) = whitelist.
	CData   string          `json:"cdata,omitempty"`   // Allows bidder to retrieve data set on its behalf in the exchange’s cookie (refer to “cdata” in Object: Response) if supported by the exchange. The string must be in base85 cookie-safe characters.
	Source  *Source         `json:"source,omitempty"`  // A Source object that provides data about the inventory source and which entity makes the final decision.
	Item    []Item          `json:"item"`              // Required. Array of Item objects (at least one) that constitute the set of goods being offered for sale.
	Package Bool            `json:"package,omitempty"` // Flag to indicate if the Exchange can verify that the items offered represent all of the items available in context (e.g., all impressions on a web page, all video spots such as pre/mid/post roll) to support road-blocking, where 0 = no, 1 = yes.
	Context json.RawMessage `json:"context,omitempty"` // Recommended. Layer-4 domain object structure that provides context for the items being offered conforming to the specification and version referenced in “openrtb.domainspec” and “openrtb.domainver”. For AdCOM v1.x, the objects allowed here all of which are optional are one of the “DistributionChannel” subtypes (i.e., “Site”, “App”, or “Dooh”), “User”, “Device”, “Regs”, and any objects subordinate to these as specified by AdCOM.
	Ext     json.RawMessage `json:"ext,omitempty"`     // Optional exchange-specific extensions.
}

// Source object carries data about the source of the transaction including the unique ID of the transaction itself, source authentication information, and the chain of custody.
type Source struct {
	TID    string          `json:"tid,omitempty"`    // Recommended. Transaction ID that must be common across all participants throughout the entire supply chain of this transaction. This also applies across all participating exchanges in a header bidding or similar publisher-centric broadcast scenario.
	DS     string          `json:"ds,omitempty"`     // Recommended. Digital signature used to authenticate this request computed by the publisher or its trusted agent from the transaction digest string "tid:digest", where ‘tid’ matches the "tid" attribute and ‘digest’ is a string composed of an immutable portion of domain objects as defined in the domain specification used for this request. Refer to Section “Inventory Authentication” for more details.
	DSGVer int8            `json:"dsgver,omitempty"` // TODO: check ads.txt doc for allowed list of versions, int8 may not be large enough to hold them all. Recommended. Format version of the digest string that was digitally signed to produce “ds”. Refer to Section “Inventory Authentication” for more details.
	Digest string          `json:"digest,omitempty"` // The full transaction digest string that was signed to produce the digital signature. Refer to Section “Inventory Authentication” for more details. NOTE: This is only intended for debugging purposes as needed. It is not intended for normal Production traffic due to the bandwidth impact.
	Cert   string          `json:"cert,omitempty"`   // Recommended. File name of the certificate (i.e., the public key) used to generate the digital signature in “ds” attribute. Refer to Section “Inventory Authentication” for more details.
	PChain string          `json:"pchain,omitempty"` // Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol. NOTE that the authentication features in this Source object combined with the “ads.txt” specification may lead to the future deprecation of this attribute.
	Ext    json.RawMessage `json:"ext,omitempty"`    //	Optional exchange-specific extensions.
}

// Item object represents a unit of goods being offered for sale either on the open market or in relation to a private marketplace deal.
// The id attribute is required since there may be multiple items being offered in the same bid request and bids must reference the specific item of interest.
// This object interfaces to Layer-4 domain objects for deeper specification of the item being offered (e.g., an impression).
type Item struct {
	ID      string          `json:"id"`                // Required. A unique identifier for this item within the context of the offer (typically starts with 1 and increments).
	Qty     int             `json:"qty,omitempty"`     // Default: 1. The number of instances (i.e., "quantity") of this item being offered (e.g., multiple identical impressions in a digital out-of-home scenario).
	Seq     *int            `json:"seq,omitempty"`     // If multiple items are offered in the same bid request, the sequence number allows for the coordinated delivery.
	Flr     float64         `json:"flr,omitempty"`     // Minimum bid price for this item expressed in CPM.
	FlrCur  string          `json:"flrcur,omitempty"`  // Default: "USD". Currency of the “flr” attribute specified using ISO-4217 alpha codes.
	Exp     int64           `json:"exp,omitempty"`     // Advisory as to the number of seconds that may elapse between auction and fulfilment.
	DT      string          `json:"dt,omitempty"`      // Date/time when the item is expected to be fulfilled (e.g. when a DOOH impression will be displayed) using the W3C-defined ISO-8601 format.
	Dlvy    int8            `json:"dlvy,omitempty"`    // Default: 0. Item (e.g., an Ad object) delivery method required, where 0 = either method, 1 = the item must be sent as part of the transaction (e.g., by value in the bid itself, fetched by URL included in the bid), and 2 = an item previously uploaded to the exchange must be referenced by its ID. Note that if an exchange does not supported prior upload, then the default of 0 is effectively the same as 1 since there can be no items to reference.
	Metric  []Metric        `json:"metric,omitempty"`  // An array of “Metric” objects.
	Deal    []Deal          `json:"deal,omitempty"`    // Array of “Deal” objects that convey special terms applicable to this item.
	Private Bool            `json:"private,omitempty"` // Default: 0 (false). Indicator of auction eligibility to seats named in “Deal” objects, where 0 (false) = all bids are accepted, 1 (true) = bids are restricted to the deals specified and the terms thereof.
	Spec    json.RawMessage `json:"spec"`              // Required. Layer-4 domain object structure that provides specifies the item being offered conforming to the specification and version referenced in “openrtb.domainspec” and “openrtb.domainver”. For AdCOM v1.x, the objects allowed here are “Placement” and any objects subordinate to these as specified by AdCOM.
	Ext     json.RawMessage `json:"ext,omitempty"`     //	Optional exchange-specific extensions.
}

// Deal object constitutes a specific deal that was struck a priori between a seller and a buyer. Its presence indicates that this item is available under the terms of that deal.
type Deal struct {
	ID       string          `json:"id"`                 // Required. A unique identifier for the deal.
	Flr      float64         `json:"flr,omitempty"`      // Minimum deal price for this item expressed in CPM.
	FlrCur   string          `json:"flrcur,omitempty"`   // Default: "USD". Currency of the “flr” attribute specified using ISO-4217 alpha codes.
	AT       int8            `json:"at,omitempty"`       // Optional override of the overall auction type of the request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in “flr” is the agreed upon deal price. Additional auction types can be defined by the exchange using 500+ values.
	WSeat    []string        `json:"wseat,omitempty"`    // Whitelist of buyer seats allowed to bid on this deal. IDs of seats and the buyer’s customers to which they refer must be coordinated between bidders and the exchange a priori. Omission implies no restrictions.
	WADomain []string        `json:"wadomain,omitempty"` // Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies no restrictions.
	Ext      json.RawMessage `json:"ext,omitempty"`      //	Optional exchange-specific extensions.
}

// Metric object is associated with an item as an array of metrics. These metrics can offer insight to assist with decisioning such as average recent viewability, click-through rate, etc. Each metric is identified by its type, reports the value of the metric, and optionally identifies the source or vendor measuring the value.
type Metric struct {
	Type   string          `json:"type"`             // Required. Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Value  float64         `json:"value"`            // Required. Number representing the value of the metric. Probabilities must be in the range 0.0 – 1.0.
	Vendor string          `json:"vendor,omitempty"` // Recommended. Source of the value using exchange curated string names which should be published to bidders a priori. If the exchange itself is the source versus a third party, "EXCHANGE" is recommended.
	Ext    json.RawMessage `json:"ext,omitempty"`    // Optional exchange-specific extensions.
}

// ----------------------------------------------------------------------------

// Response object is the bid response object under the Openrtb root.
// Its id attribute is a reflection of the bid request ID.
// The bidid attribute is an optional response tracking ID for bidders.
// If specified, it will be available for use in substitution macros placed in markup and notification URLs.
// At least one Seatbid object is required, which contains at least one Bid for an item.
// Other attributes are optional.
//
// To express a "no-bid", the most compact option is simply to return an empty response with HTTP 204.
// However, if the bidder wishes to convey a reason for not bidding, a Response object can be returned with just a reason code in the nbr attribute.
type Response struct {
	ID      string          `json:"id"`                // Required. ID of the bid request to which this is a response; must match the "request.id" attribute.
	BidID   string          `json:"bidid,omitempty"`   // Bidder generated response ID to assist with logging/tracking.
	NBR     NoBidReasonCode `json:"nbr,omitempty"`     // Reason for not bidding if applicable (see List: No-Bid Reason Codes). Note that while many exchanges prefer a simple HTTP 204 response to indicate a no-bid, responses indicating a reason code can be useful in debugging scenarios.
	Cur     string          `json:"cur,omitempty"`     // Default: "USD". Bid currency using ISO-4217 alpha codes.
	CData   string          `json:"cdata,omitempty"`   // Allows bidder to set data in the exchange’s cookie, which can be retrieved on bid requests (refer to “cdata” in Object: Request) if supported by the exchange. The string must be in base85 cookie-safe characters.
	SeatBid []SeatBid       `json:"seatbid,omitempty"` // Array of “Seatbid” objects; 1+ required if a bid is to be made.
	Ext     json.RawMessage `json:"ext,omitempty"`     //	Optional demand source specific extensions.
}

// SeatBid object.
//
// A bid response can contain multiple Seatbid objects, each on behalf of a different buyer seat and each containing one or more individual bids.
// If multiple items are presented in the request offer, the package attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if it is interested in winning any only if it can win them all as a group.
type SeatBid struct {
	Seat    string          `json:"seat,omitempty"`    // Recommended. ID of the buyer seat on whose behalf this bid is made.
	Package int8            `json:"package,omitempty"` // Default: 0. For offers with multiple items, this flag Indicates if the bidder is willing to accept wins on a subset of bids or requires the full group as a package, where 0 = individual wins accepted; 1 = package win or loss only.
	Bid     []Bid           `json:"bid"`               // Required. Array of 1+ "Bid" objects each related to an item. Multiple bids can relate to the same item.
	Ext     json.RawMessage `json:"ext,omitempty"`     // Optional demand source specific extensions.
}

// Bid object.
//
// A Seatbid object contains one or more Bid objects, each of which relates to a specific item in the bid request offer via the “item” attribute and constitutes an offer to buy that item for a given price.
type Bid struct {
	ID     string          `json:"id,omitempty"`     // Recommended. Bidder generated bid ID to assist with logging/tracking.
	Item   string          `json:"item"`             // Required. ID of the item object in the related bid request; specifically "item.id".
	Price  float64         `json:"price"`            // Required. Bid price expressed as CPM although the actual transaction is for a unit item only. Note that while the type indicates float, integer math is highly recommended when handling currencies (e.g., BigDecimal in Java).
	Deal   string          `json:"deal,omitempty"`   // Reference to a deal from the bid request if this bid pertains to a private marketplace deal; specifically “deal.id”.
	CID    string          `json:"cid,omitempty"`    // Campaign ID or other similar grouping of brand-related ads. Typically used to increase the efficiency of audit processes.
	Tactic string          `json:"tactic,omitempty"` // Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted. The specific usage and meaning of the tactic ID should be communicated between buyer and exchanges a priori.
	PURL   string          `json:"purl,omitempty"`   // Pending notice URL called by the exchange when a bid has been declared the winner within the scope of an OpenRTB compliant supply chain (i.e., there may still be non-compliant decisioning such as header bidding). Substitution macros may be included.
	BURL   string          `json:"burl,omitempty"`   // Recommended. Billing notice URL called by the exchange when a winning bid becomes billable based on exchange-specific business policy (e.g., markup rendered). Substitution macros may be included.
	LURL   string          `json:"lurl,omitempty"`   // Loss notice URL called by the exchange when a bid is known to have been lost. Substitution macros may be included. Exchange-specific policy may preclude support for loss notices or the disclosure of winning clearing prices resulting in ${OPENRTB_PRICE} macros being removed (i.e., replaced with a zero-length string).
	Exp    int64           `json:"exp,omitempty"`    // Advisory as to the number of seconds the buyer is willing to wait between auction and fulfilment.
	MID    string          `json:"mid,omitempty"`    // ID to enable media to be specified by reference if previously uploaded to the exchange rather than including it by value in the domain objects.
	Macro  []Macro         `json:"macro,omitempty"`  // Array of “Macro” objects that enable bid specific values to be substituted into markup; especially useful for previously uploaded media referenced via the “mid” attribute. Refer to Object: Macro.
	Media  json.RawMessage `json:"media,omitempty"`  // Layer-4 domain object structure that specifies the media to be presented if the bid is won conforming to the specification and version referenced in “openrtb.domainspec” and “openrtb.domainver”. For AdCOM v1.x, the objects allowed here are “Ad” and any objects subordinate thereto as specified by AdCOM.
	Ext    json.RawMessage `json:"ext,omitempty"`    // Optional demand source specific extensions.
}

// Macro object constitutes a buyer defined key/value pair used to inject dynamic values into media markup.
// While they apply to any media markup irrespective of how it is conveyed, the principle use case is for media that was uploaded to the exchange prior to the transaction (e.g., pre-registered for creative quality review) and referenced in bid.
type Macro struct {
	Key   string          `json:"key"`             // Required. Name of a buyer specific macro.
	Value string          `json:"value,omitempty"` // Value to substitute for each instance of the macro found in markup.
	Ext   json.RawMessage `json:"ext,omitempty"`   // Optional demand source specific extensions.
}
