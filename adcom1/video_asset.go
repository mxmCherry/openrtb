package adcom1

import "encoding/json"

// VideoAsset object identifies the native asset as a video asset.
// Video markup (e.g., VAST) must be either included or referenced.
type VideoAsset struct {
	// Attribute:
	//   adm
	// Type:
	//   string; required *
	// Definition:
	//   Video markup (e.g., VAST document) for the asset.
	//   * Exactly one of adm and curl is required.
	//   Including both is not recommended.
	AdM string `json:"adm,omitempty"`

	// Attribute:
	//   curl
	// Type:
	//   string; required *
	// Definition:
	//   A URL that returns the video markup (e.g., VAST document) for the asset.
	//   If this ad is matched to a placement specification, the Placement.curlx attribute indicates if this markup retrieval option is supported.
	//   * Exactly one of adm and curl is required.
	//   Including both is not recommended.
	CURL string `json:"curl,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
