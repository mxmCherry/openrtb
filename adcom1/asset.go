package adcom

import "encoding/json"

// Asset object is the container for each asset comprising a native ad.
// Each asset is of a specific type and to reflect this, one and only one of the subtype objects (i.e., title, img, video, data) must be present; all others should be omitted.
type Asset struct {
	// Attribute:
	//   id
	// Type:
	//   integer
	// Definition:
	//   The value of AssetFormat.id if this ad references a specific native placement defined by a Placement object and its structure.
	ID int `json:"id,omitempty"` // TODO: confirm type when get to AssetFormat

	// Attribute:
	//   req
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicates if the asset is required to be displayed, where 0 = no, 1 = yes.
	Req int8 `json:"req,omitempty"`

	// Attribute:
	//   title
	// Type:
	//   object; required *
	// Definition:
	//   Asset Subtype Object that indicates this is a title asset and provides additional detail as such.
	//   Refer to Object: TitleAsset.
	//   * Required if no other asset subtype object is specified.
	Title interface{} `json:"title,omitempty"`

	// Attribute:
	//   image
	// Type:
	//   object; required *
	// Definition:
	//   Asset Subtype Object that indicates this is an image asset and provides additional detail as such.
	//   Refer to Object: ImageAsset.
	//   * Required if no other asset subtype object is specified.
	Image interface{} `json:"image,omitempty"`

	// Attribute:
	//   video
	// Type:
	//   object; required *
	// Definition:
	//   Asset Subtype Object that indicates this is a video asset and provides additional detail as such.
	//   Refer to Object: VideoAsset.
	//   * Required if no other asset subtype object is specified.
	Video interface{} `json:"video,omitempty"`

	// Attribute:
	//   data
	// Type:
	//   object; required *
	// Definition:
	//   Asset Subtype Object that indicates this is a data asset and provides additional detail as such.
	//   Refer to Object: DataAsset.
	//   * Required if no other asset subtype object is specified.
	Data interface{} `json:"data,omitempty"`

	// Attribute:
	//   link
	// Type:
	//   object; required *
	// Definition:
	//   Asset Subtype Object that indicates this is a link asset and provides additional detail as such.
	//   Refer to Object: LinkAsset.
	//   * Required if no other asset subtype object is specified.
	Link interface{} `json:"link,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
