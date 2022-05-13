package adcom1

import "encoding/json"

// ExtendedIdentifier support in the OpenRTB specification allows buyers to use audience data in real-time bidding.
// The exchange should ensure that business agreements allow for the sending of this data.
// Note, it is assumed that exchanges and DSPs will collaborate with the appropriate regulatory agencies and ID vendor(s) to ensure compliance.
type ExtendedIdentifier struct {
	// Attribute:
	//   source
	// Type:
	//   string
	// Definition:
	//   Source or technology provider responsible for the set of included IDs.
	//   Expressed as a top-level domain.
	Source string `json:"source,omitempty"`

	// Attribute:
	//   uids
	// Type:
	//   object array
	// Definition:
	//   Array of extended ID UID objects from the given source. Refer to Object: Extended Identifier UIDs.
	UIDs []ExtendedIdentifierUID `json:"uids,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
