package openrtb

import "encoding/json"

// Source object carries data about the source of the transaction including the unique ID of the transaction itself, source authentication information, and the chain of custody.
//
// NOTE: Attributes ds, dsmap, cert, and digest support digitally signed bid requests as defined by the Ads.cert: Signed Bid Requests specification.
// As the Ads.cert specification is still in its BETA state, these attributes should be considered to be in a similar state.
type Source struct {
	// Attribute:
	//   tid
	// Type:
	//   string; recommended
	// Definition:
	//   Transaction ID that must be common across all participants throughout the entire supply chain of this transaction.
	//   This also applies across all participating exchanges in a header bidding or similar publisher-centric broadcast scenario.
	TID string `json:"tid,omitempty"`

	// Attribute:
	//   ts
	// Type:
	//   integer; recommended
	// Definition:
	//   Timestamp when the request originated at the beginning of the supply chain in Unix format (i.e., milliseconds since the epoch).
	//   This value must be held as immutable throughout subsequent intermediaries.
	TS int64 `json:"ts,omitempty"`

	// Attribute:
	//   ds
	// Type:
	//   string; recommended
	// Definition:
	//   Digital signature used to authenticate the origin of this request computed by the publisher or its trusted agent from a digest string composed of a set of immutable attributes found in the bid request.
	//   Refer to Section “Inventory Authentication” for more details.
	DS string `json:"ds,omitempty"`

	// Attribute:
	//   dsmap
	// Type:
	//   string
	// Definition:
	//   An ordered list of identifiers that indicates the attributes used to create the digest.
	//   This map provides the essential instructions for recreating the digest from the bid request, which is a necessary step in validating the digital signature in the ds attribute.
	//   Refer to Section “Inventory Authentication” for more details.
	DSMap string `json:"dsmap,omitempty"`

	// Attribute:
	//   cert
	// Type:
	//   string; recommended
	// Definition:
	//   File name of the certificate (i.e., the public key) used to generate the digital signature in the ds attribute.
	//   Refer to Section “Inventory Authentication” for more details.
	Cert string `json:"cert,omitempty"`

	// Attribute:
	//   digest
	// Type:
	//   string
	// Definition:
	//   The full digest string that was signed to produce the digital signature.
	//   Refer to Section “Inventory Authentication” for more details.
	//   NOTE: This is only intended for debugging purposes as needed.
	//   It is not intended for normal Production traffic due to the bandwidth impact.
	Digest string `json:"digest,omitempty"`

	// Attribute:
	//   pchain
	// Type:
	//   string
	// Definition:
	//   Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol.
	//   NOTE: Authentication features in this Source object combined with the “ads.txt” specification may lead to the deprecation of this attribute.
	PChain string `json:"pchain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional exchange-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
