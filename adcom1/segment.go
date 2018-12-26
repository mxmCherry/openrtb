package adcom

import "encoding/json"

// Segment objects are essentially key-value pairs that convey specific units of data.
// The parent Data object is a collection of such values from a given data provider.
// When in use, vendor-specific IDs should be communicated a priori among the parties.
type Segment struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   ID of the data segment specific to the data provider.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Definition:
	//   Displayable name of the data segment specific to the data provider.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   value
	// Type:
	//   string
	// Definition:
	//   String representation of the data segment value.
	Value string `json:"value,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
