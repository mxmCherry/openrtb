package adcom

import "encoding/json"

// NativeFormat object refines a display placement to be specifically a native display placement.
// It serves as the root of a structure that includes the specifications for each of the assets that comprise the native placement.
type NativeFormat struct {
	// Attribute:
	//   asset
	// Type:
	//   object array; required
	// Definition:
	//   Array of objects that specify the set of native assets and their permitted formats.
	//   Refer to Object: AssetFormat.
	Asset []interface{} `json:"asset"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
