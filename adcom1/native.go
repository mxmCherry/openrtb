package adcom

import "encoding/json"

// Native object is the root of a structure that defines a native display ad.
type Native struct {
	// Attribute:
	//   link
	// Type:
	//   object
	// Definition:
	//   Default destination link for the native ad overall; used if an asset is activated (e.g., clicked) that doesn't specify it's own destination link.
	//   Refer to Object: LinkAsset.
	Link *LinkAsset `json:"link,omitempty"`

	// Attribute:
	//   asset
	// Type:
	//   object array
	// Definition:
	//   Array of assets that comprise the native ad.
	//   Refer to Object: Asset.
	Asset []Asset `json:"asset,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
