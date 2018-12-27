package adcom

import "encoding/json"

// EventSpec object specifies a type of ad tracking event and which methods of tracking are available for it.
// This object may appear as an array for a given placement indicating various types of available tracking events.
type EventSpec struct {
	// Attribute:
	//   type
	// Type:
	//   integer; required
	// Definition:
	//   Type of supported ad tracking event.
	//   Refer to List: Event Types.
	Type EventType `json:"type,omitempty"`

	// Attribute:
	//   method
	// Type:
	//   integer array
	// Definition:
	//   Array of supported event tracking methods for this event type.
	//   Refer to List: Event Tracking Methods.
	Method []EventTrackingMethod `json:"method,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   Event tracking APIs available for use; only relevant for JavaScript method trackers.
	//   Refer to List: API Frameworks.
	API []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   jstrk
	// Type:
	//   string array
	// Definition:
	//   Array of domains, top two levels only (e.g., “tracker.com”), that constitute a restriction list of JavaScript trackers.
	//   The sense of the restrictions is determined by wjs.
	JSTrk []string `json:"jstrk,omitempty"`

	// Attribute:
	//   wjs
	// Type:
	//   integer; default 1
	// Definition:
	//   Sense of the jstrk restriction list, where 0 = block list, 1 = whitelist.
	WJS int8 `json:"wjs,omitempty"`

	// Attribute:
	//   pxtrk
	// Type:
	//   string array
	// Definition:
	//   Array of domains, top two levels only (e.g., “tracker.com”), that constitute a restriction list of pixel image trackers.
	//   The sense of the restrictions is determined by wpx.
	PxTrk []string `json:"pxtrk,omitempty"`

	// Attribute:
	//   wpx
	// Type:
	//   integer; default 1
	// Definition:
	//   Sense of the pxtrk restriction list, where 0 = block list, 1 = whitelist.
	WPx int8 `json:"wpx,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
