package adcom

import "encoding/json"

// DOOH object is used to define an ad supported digital out-of-home (DOOH) experience such as a public kiosk or digital billboard.
// As a derived class, a Dooh object inherits all DistributionChannel attributes and adds those defined below.
type DOOH struct {
	DistributionChannel

	// Attribute:
	//   venue
	// Type:
	//   integer
	// Definition:
	//   The type of out-of-home venue.
	//   Refer to List: DOOH Venue TypesList: DOOH Venue Types.
	Venue interface{} `json:"venue,omitempty"`

	// Attribute:
	//   fixed
	// Type:
	//   integer
	// Definition:
	//   Indicates whether the DOOH placement is in a fixed location (e.g., kiosk, billboard, elevator) or is movable (e.g., taxi), where 1 = fixed, 2 = movable.
	Fixed int8 `json:"fixed,omitempty"`

	// Attribute:
	//   etime
	// Type:
	//   integer
	// Definition:
	//   The exposure time in seconds per view that the creative will be displayed before refreshing to the next creative.
	ETime int64 `json:"etime,omitempty"`

	// Attribute:
	//   dpi
	// Type:
	//   integer
	// Definition:
	//   Minimum DPI for text-based creative elements to display clearly.
	DPI int `json:"dpi,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
