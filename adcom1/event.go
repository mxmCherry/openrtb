package adcom

import "encoding/json"

// Event object specifies a type of event that the advertiser or buying platform wants to track along with the information required to do so.
type Event struct {
	// Attribute:
	//   type
	// Type:
	//   integer; required
	// Definition:
	//   Type of event to track.
	//   Refer to List: Event Types.
	Type interface{} `json:"type"`

	// Attribute:
	//   method
	// Type:
	//   integer; required
	// Definition:
	//   Method of tracking requested.
	//   Refer to List: Event Tracking Methods.
	Method interface{} `json:"method"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   The APIs being used by the tracker; only relevant when the tracking method is JavaScript.
	//   Refer to List: API Frameworks.
	API []interface{} `json:"api,omitempty"`

	// Attribute:
	//   url
	// Type:
	//   string; required *
	// Definition:
	//   The URL of the tracking pixel or JavaScript tag, respectively.
	//   * Required for Image-Pixel or JavaScript methods.
	URL string `json:"url,omitempty"`

	// Attribute:
	//   cdata
	// Type:
	//   object (Map)
	// Definition:
	//   An array of key-value pairs to support vendor-specific data required for custom tracking.
	//   For example, the account number of a buyer with a tracking company might be represented as: {"acct": "123"}.
	CData map[string]string `json:"cdata,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
