package adcom1

import "encoding/json"

// Channel describes the channel an ad will be displayed on.
// A Channel is defined as the entity that curates a content library, or stream within a brand name for viewers.
// Examples are specific view selectable ‘channels’ within linear and streaming television (MTV, HGTV, CNN, BBC One, etc) or a specific stream of audio content commonly called ‘stations.’
// Name is human-readable field while domain and id can be used for reporting and targeting purposes.
type Channel struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   A unique identifier assigned by the publisher.
	//   This may not be a unique identifier across all supply sources.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Definition:
	//   Channel the content is on (e.g., a local channel like "WABC-TV").
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Definition:
	//   The primary domain of the channel (e.g., “abc7ny.com” in the case of the local channel WABC-TV).
	//   It is recommended to include the top private domain (PSL+1) for DSP targeting normalization purposes.
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
