package adcom

import "encoding/json"

// DataAsset object identifies the native asset as a data asset.
// A data asset is used for all miscellaneous elements such as brand name, ratings, stars, review count, downloads, price, counts, etc.
// It is purposefully generic to support native elements not currently contemplated by this specification.
type DataAsset struct {
	// Attribute:
	//   value
	// Type:
	//   string; required
	// Definition:
	//   A formatted string of data to be displayed (e.g., “5 stars”, “3.4 stars out of 5”, “$10”, etc.).
	Value string `json:"value"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Definition:
	//   The length of the value contents.
	//   This length should conform to recommendations provided in List: Native Data Asset Types
	Len int `json:"len,omitempty"`

	// Attribute:
	//   type
	// Type:
	//   integer
	// Definition:
	//   The type of data represented by this asset.
	//   Refer to List: Native Data Asset Types.
	Type interface{} `json:"type,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
