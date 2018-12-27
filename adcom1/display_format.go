package adcom

import "encoding/json"

// DisplayFormat object represents an allowed set of parameters for a banner display ad and often appears as an array when multiple sizes are permitted.
type DisplayFormat struct {
	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Absolute width of the creative in units specified by DisplayPlacement.unit.
	//   Note that mixing absolute and relative sizes is not recommended.
	W int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Absolute height of the creative in units specified by DisplayPlacement.unit.
	//   Note that mixing absolute and relative sizes is not recommended.
	H int `json:"h,omitempty"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Definition:
	//   Relative width of the creative when expressing size as a ratio.
	//   Note that mixing absolute and relative sizes is not recommended.
	WRatio int `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   integer
	// Definition:
	//   Relative height of the creative when expressing size as a ratio.
	//   Note that mixing absolute and relative sizes is not recommended.
	HRatio int `json:"hratio,omitempty"`

	// Attribute:
	//   expdir
	// Type:
	//   integer array
	// Definition:
	//   Directions in which the creative is permitted to expand.
	//   Refer to List: Expandable Directions.
	ExpDir []ExpandableDirection `json:"expdir,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
