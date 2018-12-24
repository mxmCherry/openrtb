package adcom

import "encoding/json"

// Banner object describes a basic banner creative.
// It is intended for display scenarios that require a simple, structured image/link pair and is more secure than allowing arbitrary HTML or JavaScript code.
type Banner struct {
	// Attribute:
	//   img
	// Type:
	//   string; required
	// Definition:
	//   A URL that will return the image.
	Img string `json:"img"`

	// Attribute:
	//   link
	// Type:
	//   object
	// Definition:
	//   Destination link if the image is activated (e.g., clicked); not applicable in some contexts (e.g., DOOH) and its inclusion does not guarantee it will be supported.
	//   Refer to Object: LinkAsset.
	Link interface{} `json:"link,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
