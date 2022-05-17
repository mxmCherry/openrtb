package openrtb2

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/v16/adcom1"
)

// 3.2.28 Object: UID
//
// This object contains a single user identifier provided as part of extended identifiers.
// The exchange should ensure that business agreements allow for the sending of this data.
type UID struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   The identifier for the user.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   atype
	// Type:
	//   object array
	// Description:
	//   Type of user agent the ID is from. It is highly recommended to set this, as
	//   many DSPs separate app-native IDs from browser-based IDs and require a type
	//   value for ID resolution. Refer to List: Agent Types in AdCOM 1.0
	AType adcom1.AgentType `json:"atype,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
