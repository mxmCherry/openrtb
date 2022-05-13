package adcom1

import "encoding/json"

// AudioPlacement object signals that the placement may be an audio placement and provides additional detail about permitted audio ads (e.g., DAAST).
type AudioPlacement struct {
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
	//   feed
	// Type:
	//   integer
	// Definition:
	//   Type of audio feed of this placement.
	//   Refer to List: Feed Types.
	Feed FeedType `json:"feed,omitempty"`

	// Attribute:
	//   nvol
	// Type:
	//   integer
	// Definition:
	//   Volume normalization mode of this placement.
	//   Refer to List: Volume Normalization Modes.
	NVol VolumeNormalizationMode `json:"nvol,omitempty"`

	// Attribute:
	//   mime
	// Type:
	//   string array; required
	// Definition:
	//   Array of supported mime types (e.g., “audio/mp4”).
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
	//   rqddurs
	// Type:
	//   integer array
	// Definition:
	//   Precise acceptable durations for video creatives in seconds. This field specifically
	//   targets the Live TV use case where non-exact ad durations would result in undesirable
	//   'dead air'. This field is mutually exclusive with mindur and maxdur; if rqddurs is
	//   specified, mindur and maxdur must not be specified and vice versa.
	RqdDurs []int64 `json:"rqddurs,omitempty"`

	// Attribute:
	//   maxext
	// Type:
	//   integer
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
	//   poddur
	// Type:
	//   integer
	// Definition:
	//   Indicates the total amount of time in seconds that advertisers may fill
	//   for a “dynamic” video ad pod, or the dynamic portion of a “hybrid” ad pod.
	//   This field is required only for the dynamic portion(s) of video ad pods.
	//   This field refers to the length of the entire ad break, whereas
	//   mindur/maxdur/rqddurs are constraints relating to the slots that make up the pod.
	PodDur int64 `json:"poddur,omitempty"`

	// Attribute:
	//   podid
	// Type:
	//   integer
	// Definition:
	//   Unique identifier indicating that an impression opportunity belongs to a
	//   video ad pod. If multiple impression opportunities within a bid request
	//   share the same podid, this indicates that those impression opportunities
	//   belong to the same video ad pod.
	PodID int64 `json:"podid,omitempty"`

	// Attribute:
	//   podseq
	// Type:
	//   integer; default 0
	// Definition:
	//   The sequence (position) of the video ad pod within a content stream.
	//   Refer to List: Pod Sequence for guidance on the use of this field.
	PodSeq PodSequence `json:"podseq,omitempty"`

	// Attribute:
	//   slotinpod
	// Type:
	//   integer; default 0
	// Definition:
	//   For video ad pods, this value indicates that the seller can guarantee delivery
	//   against the indicated slot position in the pod.
	//   Refer to List: Slot Position in Pod for guidance on the use of this field.
	SlotInPod SlotPositionInPod `json:"slotinpod,omitempty"`

	// Attribute:
	//   mincpmpersec
	// Type:
	//   float
	// Definition:
	//   Minimum CPM per second. This is a price floor for the “dynamic” portion of a
	//   video ad pod, relative to the duration of bids an advertiser may submit.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`

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
	//   overlayexpdir
	// Type:
	//   integer array
	// Definition:
	//   Directions in which the creative (video overlay) is permitted to expand.
	//   This is primarily used for non-linear videos.
	//   Refer to List: Expandable Directions.
	OverlayExpDir []ExpandableDirection `json:"overlayexpdir,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
