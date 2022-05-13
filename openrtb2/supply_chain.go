package openrtb2

import "encoding/json"

// 3.2.25 Object: SupplyChain
//
// This object is composed of a set of nodes where each node represents a specific entity that participates in the transacting of inventory.
// The entire chain of nodes from beginning to end represents all entities who are involved in the direct flow of payment for inventory.
// Detailed implementation examples can be found here: https://github.com/InteractiveAdvertisingBureau/openrtb/blob/master/supplychainobject.md.
type SupplyChain struct {

	// Attribute:
	//   complete
	// Type:
	//   integer; required
	// Description:
	//   Flag indicating whether the chain contains all nodes involved
	//   in the transaction leading back to the owner of the site, app
	//   or other medium of the inventory, where 0 = no, 1 = yes.
	Complete int8 `json:"complete"`

	// Attribute:
	//   nodes
	// Type:
	//   object array; required
	// Description:
	//   Array of SupplyChainNode objects in the order of the chain. In a
	//   complete supply chain, the first node represents the initial
	//   advertising system and seller ID involved in the transaction, i.e.
	//   the owner of the site, app, or other medium. In an incomplete
	//   supply chain, it represents the first known node. The last node
	//   epresents the entity sending this bid request.
	Nodes []SupplyChainNode `json:"nodes"`

	// Attribute:
	//   ver
	// Type:
	//   string; required
	// Description:
	//   Version of the supply chain specification in use, in the format
	//   of "major.minor". For example, for version 1.0 of the spec,
	//   use the string “1.0”.
	Ver string `json:"ver"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
