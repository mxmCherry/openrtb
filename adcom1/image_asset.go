package adcom

import "encoding/json"

// ImageAsset object identifies the native asset as a image asset.
// Image assets are use for such elements as the actual creative images and icons.
type ImageAsset struct {
	// Attribute:
	//   url
	// Type:
	//   string; required
	// Definition:
	//   A URL that returns the image for the asset.
	URL string `json:"url,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Definition:
	//   Width of the image asset in device independent pixels (DIPS).
	W int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Definition:
	//   Height of the image asset in device independent pixels (DIPS).
	H int `json:"h,omitempty"`

	// Attribute:
	//   type
	// Type:
	//   integer
	// Definition:
	//   The type of image represented by this asset.
	//   Refer to List: Native Image Asset Types.
	Type NativeImageAssetType `json:"type,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
