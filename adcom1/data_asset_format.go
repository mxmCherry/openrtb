package adcom

import "encoding/json"

// DataAssetFormat object is used to provide native asset format specifications for a data element.
// A data asset is used for all miscellaneous elements such as brand name, ratings, stars, review count, downloads, prices, etc.
// It is purposefully generic to support native elements not currently contemplated by this specification.
type DataAssetFormat struct {
	// Attribute:
	//   type
	// Type:
	//   integer; required
	// Definition:
	//   The type of data asset supported.
	//   Refer to List: Native Data Asset Types.
	Type NativeDataAssetType `json:"type,omitempty"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Definition:
	//   The maximum allowed length of the data value.
	Len int `json:"len,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
