package adcom

import "encoding/json"

// Companion object is used in video and audio placements to specify an associated or so-called companion display ad.
// Video and audio placements can specify an array of companion ads.
type Companion struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   Identifier of the companion ad; unique within this placement.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   vcm
	// Type:
	//   integer
	// Definition:
	//   Indicates the companion ad rendering mode relative to the associated video or audio ad, where 0 = concurrent, 1 = end-card.
	//   For a given placement, typically only one companion at most should be designated as an end card.
	VCm int8 `json:"vcm,omitempty"`

	// Attribute:
	//   display
	// Type:
	//   object
	// Definition:
	//   Display specification object representing the companion ad.
	//   Refer to Object: DisplayPlacement.
	Display interface{} `json:"display,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
