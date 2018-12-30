package openrtb3

import "encoding/json"

// Macro object constitutes a buyer defined key/value pair used to inject dynamic values into media markup
// While they apply to any media markup irrespective of how it is conveyed, the principle use case is for media that was uploaded to the exchange prior to the transaction (e.g., pre-registered for creative quality review) and referenced in bid
// The full form of the macro to be substituted at runtime is ${CUSTOM_KEY}, where “KEY” is the name supplied in the key attribute
// This ensures no conflict with standard OpenRTB macros.
type Macro struct {
	// Attribute:
	//   key
	// Type:
	//   string; required
	// Definition:
	//   Name of a buyer specific macro.
	Key string `json:"key"`

	// Attribute:
	//   value
	// Type:
	//   string
	// Definition:
	//   Value to substitute for each instance of the macro found in markup.
	Value string `json:"value,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional demand source specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
