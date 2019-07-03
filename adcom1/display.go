package adcom1

import "encoding/json"

// Display object provides additional detail about an ad specifically for display ads.
// There are multiple attributes for specifying creative details: banner for simple banner images native for native ads, adm for including general markup, and curl for referencing general markup via URL.
// In any given Display object, only one of these attributes should be used to avoid confusion.
// To the extent feasible, structured objects should be favored over general markup for quality and safety issues.
type Display struct {
	// Attribute:
	//   mime
	// Type:
	//   string
	// Definition:
	//   Mime type of the ad (e.g., “image/jpeg”).
	MIME string `json:"mime,omitempty"`

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
	//   Subtype of display creative.
	//   Refer to List: Creative Subtypes - Display.
	CType DisplayCreativeSubtype `json:"ctype,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Absolute width of the creative in device independent pixels (DIPS), typically for non-native ads.
	//   Note that mixing absolute and relative sizes is not recommended.
	W int64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Absolute height of the creative in device independent pixels (DIPS), typically for non-native ads.
	//   Note that mixing absolute and relative sizes is not recommended.
	H int64 `json:"h,omitempty"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Definition:
	//   Relative width of the creative when expressing size as a ratio, typically for non-native ads.
	//   Note that mixing absolute and relative sizes is not recommended.
	// Dev note:
	//   This is kept as `int8` because ratio values are expected to be quite small (like 16:9).
	WRatio int8 `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   integer
	// Definition:
	//   Relative height of the creative when expressing size as a ratio, typically for non-native ads.
	//   Note that mixing absolute and relative sizes is not recommended.
	// Dev note:
	//   This is kept as `int8` because ratio values are expected to be quite small (like 16:9).
	HRatio int8 `json:"hratio,omitempty"`

	// Attribute:
	//   priv
	// Type:
	//   string
	// Definition:
	//   URL of a page informing the user about a buyer's targeting activity.
	Priv string `json:"priv,omitempty"`

	// Attribute:
	//   adm
	// Type:
	//   string
	// Definition:
	//   General display markup (e.g., HTML, AMPHTML) if not using a structured alternative (e.g., banner, native).
	//   Note that including both adm and curl is not recommended.
	AdM string `json:"adm,omitempty"`

	// Attribute:
	//   curl
	// Type:
	//   string
	// Definition:
	//   Optional means of retrieving display markup by reference; a URL that can return HTML, AMPHTML, or a collection native Asset object and their subordinates).
	//   If this ad is matched to a Placement specification, the Placement.curlx attribute indicates if this markup retrieval option is supported.
	//   Note that including both adm and curl is not recommended.
	CURL string `json:"curl,omitempty"`

	// Attribute:
	//   banner
	// Type:
	//   object
	// Definition:
	//   Structured banner image object, recommended for simple banner creatives.
	//   Refer to Object: Banner.
	Banner *Banner `json:"banner,omitempty"`

	// Attribute:
	//   native
	// Type:
	//   object
	// Definition:
	//   Structured native object, recommended for native ads.
	//   Refer to Object: Native.
	Native *Native `json:"native,omitempty"`

	// Attribute:
	//   event
	// Type:
	//   object array
	// Definition:
	//   Array of events that the advertiser or buying platform wants to track.
	//   Refer to Object: Event.
	Event []Event `json:"event,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
