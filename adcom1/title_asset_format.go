package adcom1

import "encoding/json"

// TitleAssetFormat object is used to provide native asset format specifications for a title element.
// Title elements are simple strings.
type TitleAssetFormat struct {
	// Attribute:
	//   len
	// Type:
	//   integer; required
	// Definition:
	//   The maximum allowed length of the title value.
	//   Recommended lengths are 25, 90, or 140.
	Len int64 `json:"len,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
