package adcom1

import "encoding/json"

// LinkAsset object identifies the native asset as a link asset and is used to define navigation for call-to-action or other activations (i.e., clicks)
// A link asset can be independent or associated with the overall native ad (i.e., a default for all assets).
type LinkAsset struct {
	// Attribute:
	//   url
	// Type:
	//   string; required
	// Definition:
	//   Landing URL of the clickable link.
	URL string `json:"url,omitempty"`

	// Attribute:
	//   urlfb
	// Type:
	//   string
	// Definition:
	//   Fallback URL for deep-link to be used if the URL specified in url is not supported by the device.
	URLFB string `json:"urlfb,omitempty"`

	// Attribute:
	//   trkr
	// Type:
	//   string array
	// Definition:
	//   List of third-party tracker URLs to be fired on click.
	Trkr []string `json:"trkr,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
