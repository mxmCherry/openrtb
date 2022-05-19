package openrtb2

import "encoding/json"

// 3.2.24 Object: Channel
//
// This object describes the channel an ad will be displayed on.
// A Channel is defined as the entity that curates a content library, or stream within a brand name for viewers.
// Examples are specific view selectable ‘channels’ within linear and streaming television (MTV, HGTV, CNN, BBC One, etc) or a specific stream of audio content commonly called ‘stations.’
// Name is a human-readable field while domain and id can be used for reporting and targeting purposes.
// See 7.6 for further examples.
type Channel struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   A unique identifier assigned by the publisher. This may not be
	//   a unique identifier across all supply sources.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Channel the content is on (e.g., a local channel like “WABC-TV").
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   The primary domain of the channel (e.g. "abc7ny.com" in the
	//   case of the local channel WABC-TV). It is recommended to
	//   include the top private domain (PSL+1) for DSP targeting
	//   normalization purposes.
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
