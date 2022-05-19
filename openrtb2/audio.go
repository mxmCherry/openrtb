package openrtb2

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/v16/adcom1"
)

// 3.2.8 Object: Audio
//
// This object represents an audio type impression.
// Many of the fields are non-essential for minimally viable transactions, but are included to offer fine control when needed.
// Audio in OpenRTB generally assumes compliance with the DAAST standard.
// As such, the notion of companion ads is supported by optionally including an array of Banner objects (refer to the Banner object in Section 3.2.6) that define these companion ads.
//
// The presence of a Audio as a subordinate of the Imp object indicates that this impression is offered as an audio type impression.
// At the publisher’s discretion, that same impression may also be offered as banner, video, and/or native by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.
type Audio struct {

	// Attribute:
	//   mimes
	// Type:
	// string array; required
	// Description:
	//   Content MIME types supported (e.g., “audio/mp4”).
	MIMEs []string `json:"mimes"`

	// Attribute:
	//   minduration
	// Type:
	//   integer; default 0; recommended
	// Description:
	//   Minimum audio ad duration in seconds.
	MinDuration int64 `json:"minduration,omitempty"`

	// Attribute:
	//   maxduration
	// Type:
	//   integer; recommended
	// Description:
	//   Maximum audio ad duration in seconds.
	MaxDuration int64 `json:"maxduration,omitempty"`

	// Attribute:
	//   poddur
	// Type:
	//   integer; recommended
	// Description:
	//   Indicates the total amount of time that advertisers may fill for a
	//   "dynamic" audio ad pod, or the dynamic portion of a "hybrid"
	//   ad pod. This field is required only for the dynamic portion(s) of
	//   audio ad pods. This field refers to the length of the entire ad
	//   break, whereas minduration/maxduration/rqddurs are
	//   constraints relating to the slots that make up the pod.
	PodDur int64 `json:"poddur,omitempty"`

	// Attribute:
	//   protocols
	// Type:
	//   integer array; recommended
	// Description:
	//   Array of supported audio protocols. Refer to List: Creative
	//   Subtypes - Audio/Video in AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only protocols 1..10.
	Protocols []adcom1.MediaCreativeSubtype `json:"protocols,omitempty"`

	// Attribute:
	//   startdelay
	// Type:
	//   integer; recommended
	// Description:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or
	//   post-roll ad placements. Refer to List: Start Delay Modes in
	//   AdCOM 1.0.
	StartDelay *adcom1.StartDelay `json:"startdelay,omitempty"`

	// Attribute:
	//   rqddurs
	// Type:
	//   integer array
	// Description:
	//   Precise acceptable durations for audio creatives in seconds. This
	//   field specifically targets the live audio/radio use case where
	//   non-exact ad durations would result in undesirable ‘dead air’.
	//   This field is mutually exclusive with minduration and
	//   maxduration; if rqddurs is specified, minduration and
	//   maxduration must not be specified and vice versa.
	RqdDurs []int64 `json:"rqddurs,omitempty"`

	// Attribute:
	//   podid
	// Type:
	//   integer
	// Description:
	//   Unique identifier indicating that an impression opportunity
	//   belongs to an audioad pod. If multiple impression opportunities
	//   within a bid request share the same podid, this indicates that
	//   those impression opportunities belong to the same audio ad
	//   pod.
	PodID int64 `json:"podid,omitempty"`

	// Attribute:
	//   podid
	// Type:
	//   integer; default 0
	// Description:
	//   The sequence (position) of the audio ad pod within a
	//   content stream. Refer to List: Pod Sequence in AdCOM 1.0
	//   for guidance on the use of this field.
	PodSeq adcom1.PodSequence `json:"podseq,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer; default 0; DEPRECATED
	// Description:
	//   If multiple ad impressions are offered in the same bid request,
	//   the sequence number will allow for the coordinated delivery
	//   of multiple creatives.
	Sequence int64 `json:"sequence,omitempty"`

	// Attribute:
	//   slotinpod
	// Type:
	//   integer; default 0
	// Description:
	//   For audio ad pods, this value indicates that the seller can
	//   guarantee delivery against the indicated sequence. Refer to
	//   List: Slot Position in Pod in AdCOM 1.0 for guidance on the
	//   use of this field.
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`

	// Attribute:
	//   mincpmpersec
	// Type:
	//   float
	// Description:
	//   Minimum CPM per second. This is a price floor for the
	//   "dynamic" portion of an audio ad pod, relative to the duration
	//   of bids an advertiser may submit.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List: Creative Attributes in
	//   AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only attributes with IDs 1..17.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   maxextended
	// Type:
	//   integer
	// Description:
	//   Maximum extended ad duration if extension is allowed. If
	//   blank or 0, extension is not allowed. If -1, extension is
	//   allowed, and there is no time limit imposed. If greater than 0,
	//   then the value represents the number of seconds of extended
	//   play supported beyond the maxduration value.
	MaxExtended int64 `json:"maxextended,omitempty"`

	// Attribute:
	//   minbitrate
	// Type:
	//   integer
	// Description:
	//   Minimum bit rate in Kbps.
	MinBitrate int64 `json:"minbitrate,omitempty"`

	// Attribute:
	//   maxbitrate
	// Type:
	//   integer
	// Description:
	//   Maximum bit rate in Kbps.
	MaxBitrate int64 `json:"maxbitrate,omitempty"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Description:
	//   Supported delivery methods (e.g., streaming, progressive). If
	//   none specified, assume all are supported. Refer to List: Delivery
	//   Methods in AdCOM 1.0.
	Delivery []adcom1.DeliveryMethod `json:"delivery,omitempty"`

	// Attribute:
	//   companionad
	// Type:
	//   object array
	// Description:
	//   Array of Banner objects (Section 3.2.6) if companion ads are
	//   available.
	CompanionAd []Banner `json:"companionad,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List: API Frameworks in AdCOM 1.0. If an API is not explicitly
	//   listed, it is assumed not to be supported.
	// Note:
	//   OpenRTB <=2.5 defined only frameworks 1..6.
	API []adcom1.APIFramework `json:"api,omitempty"`

	// Attribute:
	//   companiontype
	// Type:
	//   integer array
	// Description:
	//   Supported companion ad types. Refer to List: Companion
	//   Types in AdCOM 1.0. Recommended if companion Banner
	//   objects are included via the companionad array.
	CompanionType []adcom1.CompanionType `json:"companiontype,omitempty"`

	// Attribute:
	//   maxseq
	// Type:
	//   integer
	// Description:
	//   The maximum number of ads that can be played in an ad pod.
	MaxSeq int64 `json:"maxseq,omitempty"`

	// Attribute:
	//   feed
	// Type:
	//   integer
	// Description:
	//   Type of audio feed. Refer to List: Feed Types in AdCOM 1.0.
	Feed adcom1.FeedType `json:"feed,omitempty"`

	// Attribute:
	//   stitched
	// Type:
	//   integer
	// Description:
	//   Indicates if the ad is stitched with audio content or delivered
	//   independently, where 0 = no, 1 = yes.
	Stitched int8 `json:"stitched,omitempty"`

	// Attribute:
	//   nvol
	// Type:
	//   integer
	// Description:
	//   Volume normalization mode. Refer to List: Volume
	//   Normalization Modes in AdCOM 1.0.
	NVol *adcom1.VolumeNormalizationMode `json:"nvol,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
