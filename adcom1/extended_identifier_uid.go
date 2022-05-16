package adcom1

import "encoding/json"

type ExtendedIdentifierUID struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   Cookie or platform-native identifier.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   atype
	// Type:
	//   integer
	// Definition:
	//   Type of user agent the match is from.
	//   It is highly recommended to set this, as many DSPs separate app-native IDs from browser-based IDs and require a type value for ID resolution.
	//   Refer to List: Agent Types.
	AType AgentType `json:"atype,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
