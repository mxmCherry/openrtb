package adcom1

import "encoding/json"

// Audit objects represents the outcome of some form of review of the ad.
// This is typical, for example, when scanning for malware or otherwise performing ad quality reviews.
type Audit struct {
	// Attribute:
	//   status
	// Type:
	//   integer
	// Definition:
	//   The audit status of the ad.
	//   Refer to List: Audit Status Codes.
	Status AuditStatus `json:"status,omitempty"`

	// Attribute:
	//   feedback
	// Type:
	//   string array
	// Definition:
	//   One or more human-readable explanations as to reasons for rejection or any changes to fields for ad quality reasons (e.g., adomain, cat, attr, etc.).
	Feedback []string `json:"feedback,omitempty"`

	// Attribute:
	//   init
	// Type:
	//   integer
	// Definition:
	//   Timestamp of the original instantiation of this object in Unix format (i.e., milliseconds since the epoch).
	Init int64 `json:"init,omitempty"`

	// Attribute:
	//   lastmod
	// Type:
	//   integer
	// Definition:
	//   Timestamp of most recent modification to this object in Unix format (i.e., milliseconds since the epoch).
	LastMod int64 `json:"lastmod,omitempty"`

	// Attribute:
	//   corr
	// Type:
	//   object
	// Definition:
	//   Correction object wherein the auditor can specify changes to attributes of the Ad object or its children they believe to be proper.
	//   For example, if the original Ad indicated a category of “IAB3”, but the auditor deems the correct category to be “IAB13”, then corr could include a sparse Ad object including just the cat array indicating “IAB13”.
	Corr *Ad `json:"corr,omitempty"` // TODO: probably, this won't work due to "omitempty" stuff. Probably, will need an all-pointer Ad equivalent.

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
