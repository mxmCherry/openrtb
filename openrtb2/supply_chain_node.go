package openrtb2

import "encoding/json"

// 3.2.26 Object: SupplyChainNode
//
// This object is associated with a SupplyChain object as an array of nodes.
// These nodes define the identity of an entity participating in the supply chain of a bid request.
// Detailed implementation examples can be found here: https://github.com/InteractiveAdvertisingBureau/openrtb/blob/master/supplychainobject.md.
type SupplyChainNode struct {

	// Attribute:
	//   asi
	// Type:
	//   string; required
	// Description:
	//   The canonical domain name of the SSP, Exchange, Header
	//   Wrapper, etc system that bidders connect to. This may be
	//   the operational domain of the system, if that is different than
	//   the parent corporate domain, to facilitate WHOIS and
	//   reverse IP lookups to establish clear ownership of the
	//   delegate system. This should be the same value as used to
	//   identify sellers in an ads.txt file if one exists
	ASI string `json:"asi"`

	// Attribute:
	//   sid
	// Type:
	//   string; required
	// Description:
	//   The identifier associated with the seller or reseller account
	//   within the advertising system. This must contain the same value
	//   used in transactions (i.e. OpenRTB bid requests) in the field
	//   specified by the SSP/exchange. Typically, in OpenRTB, this is
	//   publisher.id. For OpenDirect it is typically the publisher’s
	//   organization ID.Should be limited to 64 characters in length.
	SID string `json:"sid"`

	// Attribute:
	//   rid
	// Type:
	//   string
	// Description:
	//   The OpenRTB RequestId of the request as issued by this seller.
	RID string `json:"rid,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   The name of the company (the legal entity) that is paid for
	//   inventory transacted under the given seller_ID. This value is
	//   optional and should NOT be included if it exists in the
	//   advertising system’s sellers.json file.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//  The business domain name of the entity represented by this
	//  node. This value is optional and should NOT be included if it
	//  exists in the advertising system’s sellers.json file.
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   hp
	// Type:
	//   integer; required
	// Description:
	//   Indicates whether this node will be involved in the flow of
	//   payment for the inventory. When set to 1, the advertising
	//   system in the asi field pays the seller in the sid field, who is
	//   responsible for paying the previous node in the chain. When
	//   set to 0, this node is not involved in the flow of payment for
	//   the inventory. For version 1.0 of SupplyChain, this property
	//   should always be 1. Implementers should ensure that they
	//   propagate this field onwards when constructing SupplyChain
	//   objects in bid requests sent to a downstream advertising
	//   system.
	HP int8 `json:"hp"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
