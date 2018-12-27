package adcom

import "encoding/json"

// Placement object represents the properties of a placement and the characteristics of ads permitted to be rendering within them.
// Placements of all types begin with this object as their root.
// It contains one or more subtype objects (i.e., display, video, audio) that define the kinds of media permitted as well as media specific placement behaviors.
//
// The other attributes in this object apply to all aspects and substructures of the placement (i.e., the same language, secure status, etc. apply to all media types and native assets as applicable).
type Placement struct {
	// Attribute:
	//   tagid
	// Type:
	//   string
	// Definition:
	//   Identifier for specific ad placement or ad tag; unique within a distribution channel.
	TagID string `json:"tagid,omitempty"`

	// Attribute:
	//   ssai
	// Type:
	//   Integer; default 0
	// Definition:
	//   Indicates if server-side ad insertion (e.g., stitching an ad into an audio or video stream) is in use and the impact of this on asset and tracker retrieval, where 0 = status unknown, 1 = all client-side (i.e., not server-side), 2 = assets stitched server-side but tracking pixels fired client-side, 3 = all server-side.
	SSAI int8 `json:"ssai,omitempty"`

	// Attribute:
	//   sdk
	// Type:
	//   string
	// Definition:
	//   Name of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video, audio, or mobile); used by some ad servers to customize ad code by partner.
	SDK string `json:"sdk,omitempty"`

	// Attribute:
	//   sdkver
	// Type:
	//   string
	// Definition:
	//   Version of the SDK specified in the sdk attribute.
	SDKVer string `json:"sdkver,omitempty"`

	// Attribute:
	//   reward
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicates if this is a rewarded placement, where 0 = no, 1 = yes.
	Reward int8 `json:"reward,omitempty"`

	// Attribute:
	//   wlang
	// Type:
	//   string array
	// Definition:
	//   Whitelist of permitted languages of the creative using ISO-639-1-alpha-2.
	//   In practice, vendors using this object may elect an alternate standard (e.g., BCP-47) in which case this must be communicated a priori.
	//   Omission of this attribute indicates there are no restrictions.
	WLang []string `json:"wlang,omitempty"`

	// Attribute:
	//   secure
	// Type:
	//   integer
	// Definition:
	//   Flag to indicate if the creative is secure (i.e., uses HTTPS for all assets and markup), where 0 = no, 1 = yes.
	//   There is no default and thus if omitted, the secure state is unknown.
	//   However, as a practical matter, the safe assumption is to treat unknown as non-secure.
	Secure int8 `json:"secure,omitempty"`

	// Attribute:
	//   admx
	// Type:
	//   integer
	// Definition:
	//   Indicates if including markup is supported (i.e., the various adm attributes throughout the Placement structure), where 0 = no, 1 = yes.
	AdMX int8 `json:"admx,omitempty"`

	// Attribute:
	//   curlx
	// Type:
	//   integer
	// Definition:
	//   Indicates if retrieving markup via URL reference is supported (i.e., the various curl attributes throughout the Placement structure), where 0 = no, 1 = yes.
	CURLX int8 `json:"curlx,omitempty"`

	// Attribute:
	//   display
	// Type:
	//   object; required *
	// Definition:
	//   Placement Subtype Object that indicates that this may be a display placement and provides additional detail related thereto.
	//   Refer to Object: DisplayPlacement.
	//   * At least one placement subtype object is required.
	Display interface{} `json:"display,omitempty"`

	// Attribute:
	//   video
	// Type:
	//   object; required *
	// Definition:
	//   Placement Subtype Object that indicates that this may be a video placement and provides additional detail related thereto.
	//   Refer to Object: VideoPlacement.
	//   * At least one placement subtype object is required.
	Video interface{} `json:"video,omitempty"`

	// Attribute:
	//   audio
	// Type:
	//   object; required *
	// Definition:
	//   Placement Subtype Object that indicates that this may be an audio placement and provides additional detail related thereto.
	//   Refer to Object: AudioPlacement.
	//   * At least one placement subtype object is required.
	Audio interface{} `json:"audio,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
