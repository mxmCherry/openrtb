package adcom1

import "encoding/json"

// Regs object contains any known legal, governmental, or industry regulations that are in effect.
type Regs struct {
	// Attribute:
	//   coppa
	// Type:
	//   integer
	// Definition:
	//   Flag indicating if COPPA regulations apply, where 0 = no, 1 = yes.
	//   The Children's Online Privacy Protection Act (COPPA) was established by the U.S. Federal Trade Commission.
	COPPA int8 `json:"coppa,omitempty"`

	// Attribute:
	//   gdpr
	// Type:
	//   integer
	// Definition:
	//   Flag indicating if GDPR regulations apply, where 0 = no, 1 = yes.
	//   The General Data Protection Regulation (GDPR) is a regulation of the European Union.
	GDPR int8 `json:"gdpr,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
