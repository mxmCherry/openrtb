package adcom1

import "encoding/json"

// TitleAsset object identifies the native asset as a title asset, which is essentially just a plain text string with specified length.
type TitleAsset struct {
	// Attribute:
	//   text
	// Type:
	//   string; required
	// Definition:
	//   The text content of the text element.
	Text string `json:"text,omitempty"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Definition:
	//   The length of the contents of the text attribute.
	Len int64 `json:"len,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
