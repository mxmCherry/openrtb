package openrtb2

import "encoding/json"

// 3.2.23 Object: Network
//
// This object describes the network an ad will be displayed on.
// A Network is defined as the parent entity of the Channel object’s entity for the purposes of organizing Channels.
// Examples are companies that own and/or license a collection of content channels (Viacom, Discovery, CBS, WarnerMedia, Turner and others), or studio that creates such content and self-distributes content.
// Name is a human-readable field while domain and id can be used for reporting and targeting purposes.
// See 7.6 for further examples.
type Network struct {

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
	//   Network the content is on (e.g., a TV network like “ABC").
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   The primary domain of the network (e.g. "abc.com" in the
	//   case of the network ABC). It is recommended to include the
	//   top private domain (PSL+1) for DSP targeting normalization
	//   purposes.
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
