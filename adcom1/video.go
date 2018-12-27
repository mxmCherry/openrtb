package adcom

import "encoding/json"

// Video object provides additional detail about an ad specifically for video ads.
type Video struct {
	// Attribute:
	//   mime
	// Type:
	//   string array
	// Definition:
	//   Mime type(s) of the ad creative(s) (e.g., “video/mp4”).
	MIME []string `json:"mime,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   API required by the ad if applicable.
	//   Refer to List: API Frameworks.
	API []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   ctype
	// Type:
	//   integer
	// Definition:
	//   Subtype of video creative.
	//   Refer to List: Creative Subtypes - Audio/Video.
	CType CreativeSubtypeAV `json:"ctype,omitempty"`

	// Attribute:
	//   dur
	// Type:
	//   integer
	// Definition:
	//   Duration of the video creative in seconds.
	Dur int64 `json:"dur,omitempty"`

	// Attribute:
	//   adm
	// Type:
	//   string
	// Definition:
	//   Video markup (e.g., VAST).
	//   Note that including both adm and curl is not recommended.
	AdM string `json:"adm,omitempty"`

	// Attribute:
	//   curl
	// Type:
	//   string
	// Definition:
	//   Optional means of retrieving markup by reference; a URL that returns video markup (e.g., VAST).
	//   If this ad is matched to a Placement specification, the Placement.curlx attribute indicates if this markup retrieval option is supported.
	//   Note that including both adm and curl is not recommended.
	CURL string `json:"curl,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
