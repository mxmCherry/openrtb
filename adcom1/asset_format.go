package adcom

import "encoding/json"

// AssetFormat object represents the permitted specifications of a single asset of a native ad.
// Along with its own attributes, exactly one of the asset subtype objects must be included.
// All others must be omitted.
type AssetFormat struct {
	// Attribute:
	//   id
	// Type:
	//   integer; required
	// Definition:
	//   Asset ID, unique within the scope of this placement specification.
	ID int `json:"id"`

	// Attribute:
	//   req
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicator of whether or not this asset is required, where 0 = no, 1 = yes.
	Req int8 `json:"req,omitempty"`

	// Attribute:
	//   title
	// Type:
	//   object; required *
	// Definition:
	//   Asset Format Subtype Object that indicates this is specifying a title asset and provides additional detail as such.
	//   Refer to Object: TitleAssetFormat.
	//   * Required if no other asset format subtype object is specified.
	Title interface{} `json:"title,omitempty"`

	// Attribute:
	//   img
	// Type:
	//   object; required *
	// Definition:
	//   Asset Format Subtype Object that indicates this is specifying an image asset and provides additional detail as such.
	//   Refer to Object: ImageAssetFormat.
	//   * Required if no other asset format subtype object is specified.
	Img interface{} `json:"img,omitempty"`

	// Attribute:
	//   video
	// Type:
	//   object; required *
	// Definition:
	//   Asset Format Subtype Object, which leverages the VideoPlacement object, that indicates this is specifying a video asset and provides additional detail as such.
	//   Refer to Object: VideoPlacement.
	//   * Required if no other asset format subtype object is specified.
	Video interface{} `json:"video,omitempty"`

	// Attribute:
	//   data
	// Type:
	//   object; required *
	// Definition:
	//   Asset Format Subtype Object that indicates this is specifying a data asset and provides additional detail as such.
	//   Refer to Object: DataAssetFormat.
	//   * Required if no other asset format subtype object is specified.
	Data interface{} `json:"data,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
