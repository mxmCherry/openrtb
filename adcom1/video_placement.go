package adcom1

import "encoding/json"

// VideoPlacement object signals that the placement may be a video placement and provides additional detail about permitted video ads (e.g., VAST).
type VideoPlacement struct {
	// Attribute:
	//   ptype
	// Type:
	//   integer
	// Definition:
	//   Placement subtype.
	//   Refer to List: Placement Subtypes - Video.
	PType VideoPlacementSubtype `json:"ptype,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Definition:
	//   Placement position on screen.
	//   Refer to List: Placement Positions.
	Pos PlacementPosition `json:"pos,omitempty"`

	// Attribute:
	//   delay
	// Type:
	//   integer
	// Definition:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll placements.
	//   For additional generic values, refer to List: Start Delay Modes.
	Delay int64 `json:"delay,omitempty"`

	// Attribute:
	//   skip
	// Type:
	//   integer
	// Definition:
	//   Indicates if the placement imposes ad skippability, where 0 = no, 1 = yes.
	Skip int8 `json:"skip,omitempty"`

	// Attribute:
	//   skipmin
	// Type:
	//   integer; default 0
	// Definition:
	//   The placement allows creatives of total duration greater than this number of seconds to be skipped; only applicable if the ad is skippable.
	SkipMin int64 `json:"skipmin,omitempty"`

	// Attribute:
	//   skipafter
	// Type:
	//   integer; default 0
	// Definition:
	//   Number of seconds a creative must play before the placement enables skipping; only applicable if the ad is skippable.
	SkipAfter int64 `json:"skipafter,omitempty"`

	// Attribute:
	//   playmethod
	// Type:
	//   integer
	// Definition:
	//   Playback method in use for this placement.
	//   Refer to List: Playback Methods.
	PlayMethod PlaybackMethod `json:"playmethod,omitempty"`

	// Attribute:
	//   playend
	// Type:
	//   integer
	// Definition:
	//   The event that causes playback to end for this placement.
	//   Refer to List: Playback Cessation Modes.
	PlayEnd PlaybackCessationMode `json:"playend,omitempty"`

	// Attribute:
	//   clktype
	// Type:
	//   integer
	// Definition:
	//   Indicates the click type of the placement.
	//   Refer to List: Click Types.
	ClkType ClickType `json:"clktype,omitempty"`

	// Attribute:
	//   mime
	// Type:
	//   string array; required
	// Definition:
	//   Array of supported mime types (e.g., “video/mp4”).
	//   If omitted, all types are assumed.
	MIME []string `json:"mime,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   List of supported APIs for this placement.
	//   If an API is not explicitly listed, it is assumed to be unsupported.
	//   Refer to List: API Frameworks.
	API []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   ctype
	// Type:
	//   integer array
	// Definition:
	//   Creative subtypes permitted for this placement.
	//   Refer to List: Creative Subtypes - Audio/Video.
	CType []MediaCreativeSubtype `json:"ctype,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Width of the placement in units specified by unit.
	W int64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Height of the placement in units specified by unit.
	H int64 `json:"h,omitempty"`

	// Attribute:
	//   unit
	// Type:
	//   integer; default 1
	// Definition:
	//   Units of size used for w and h attributes.
	//   Refer to List: Size Units.
	Unit SizeUnit `json:"unit,omitempty"`

	// Attribute:
	//   mindur
	// Type:
	//   integer
	// Definition:
	//   Minimum creative duration in seconds.
	MinDur int64 `json:"mindur,omitempty"`

	// Attribute:
	//   maxdur
	// Type:
	//   integer
	// Definition:
	//   Maximum creative duration in seconds.
	MaxDur int64 `json:"maxdur,omitempty"`

	// Attribute:
	//   maxext
	// Type:
	//   integer; default 0
	// Definition:
	//   Maximum extended creative duration if extension is allowed.
	//   If 0, extension is not allowed.
	//   If -1, extension is allowed and there is no time limit imposed.
	//   If greater than 0, then the value represents the number of seconds of extended play supported beyond the maxdur value.
	MaxExt int64 `json:"maxext,omitempty"`

	// Attribute:
	//   minbitr
	// Type:
	//   integer
	// Definition:
	//   Minimum bit rate of the creative in Kbps.
	MinBitR int64 `json:"minbitr,omitempty"`

	// Attribute:
	//   maxbitr
	// Type:
	//   integer
	// Definition:
	//   Maximum bit rate of the creative in Kbps.
	MaxBitR int64 `json:"maxbitr,omitempty"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Definition:
	//   Array of supported creative delivery methods.
	//   If omitted, all can be assumed.
	//   Refer to List: Delivery Methods.
	Delivery []DeliveryMethod `json:"delivery,omitempty"`

	// Attribute:
	//   maxseq
	// Type:
	//   integer
	// Definition:
	//   The maximum number of ads that can be played in an ad pod.
	MaxSeq int64 `json:"maxseq,omitempty"`

	// Attribute:
	//   linear
	// Type:
	//   integer
	// Definition:
	//   Indicates if the creative must be linear, nonlinear, etc.
	//   If none specified, no restrictions are assumed.
	//   Refer to List: Linearity Modes.
	Linear LinearityMode `json:"linear,omitempty"`

	// Attribute:
	//   boxing
	// Type:
	//   integer; default 1
	// Definition:
	//   Indicates if letterboxing of 4:3 creatives into a 16:9 window is allowed, where 0 = no, 1 = yes.
	Boxing int8 `json:"boxing,omitempty"`

	// Attribute:
	//   comp
	// Type:
	//   object array
	// Definition:
	//   Array of objects indicating that companion ads are available and providing the specifications thereof.
	//   Refer to Object: Companion.
	Comp []Companion `json:"comp,omitempty"`

	// Attribute:
	//   comptype
	// Type:
	//   integer array
	// Definition:
	//   Supported companion ad types; recommended if companion ads are specified in comp.
	//   Refer to List: Companion Types.
	CompType []CompanionType `json:"comptype,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
