package rtb

// Maximum extended video ad duration if extension is allowed.
// If blank or 0, extension is not allowed. If -1, extension is
// allowed, and there is no time limit imposed. If greater than 0,
// then the value represents the number of seconds of extended
// play supported beyond the maxduration value.
const (
	VideoMaxExtendedUnlimited  int64 = -1 // If -1, extension is allowed, and there is no time limit imposed
	VideoMaxExtendedNotAllowed int64 =  0 // If blank or 0, extension is not allowed
)

// Indicates if letter-boxing of 4:3 content into a 16:9 window is
// allowed, where 0 = no, 1 = yes.
const (
	VideoBoxingAllowedNo  uint8 = 0 // 0 = no
	VideoBoxingAllowedYes uint8 = 1 // 1 = yes
)

// 3.2.4 Object: Video
//   This object represents an in-stream video impression. Many of the fields are non-essential for minimally
//   viable transactions, but are included to offer fine control when needed. Video in OpenRTB generally
//   assumes compliance with the VAST standard. As such, the notion of companion ads is supported by
//   optionally including an array of Banner objects (refer to the Banner object in Section 3.2.3) that define
//   these companion ads.
//
//   The presence of a Video as a subordinate of the Imp object indicates that this impression is offered as a
//   video type impression. At the publisher’s discretion, that same impression may also be offered as
//   banner and/or native by also including as Imp subordinates the Banner and/or Native objects,
//   respectively. However, any given bid for the impression must conform to one of the offered types.
type Video struct {

	// Attribute:
	//   mimes
	// Type:
	//   string array
	// Description:
	//   Content MIME types supported. Popular MIME types may include “video/x-ms-wmv” for Windows Media and
	//   “video/x-flv” for Flash Video.
	MIMEs []string `json:"mimes"`

	// Attribute:
	//   minduration
	// Type:
	//   integer; recommended
	// Description:
	//   Minimum video ad duration in seconds.
	MinDuration int64 `json:"minduration"`

	// Attribute:
	//   maxduration
	// Type:
	//   integer; recommended
	// Description:
	//   Maximum video ad duration in seconds.
	MaxDuration int64 `json:"maxduration"`

	// Attribute:
	//   protocol
	// Type:
	//   integer array; DEPRECATED
	// Description:
	//   NOTE: Use of protocols instead is highly recommended.
	//   Supported video bid response protocol. Refer to List 5.8. At
	//   least one supported protocol must be specified in either the
	//   protocol or protocols attribute.
	Protocol []int64 `json:"protocol"`

	// Attribute:
	//   protocols
	// Type:
	//   integer array; recommended
	// Description:
	//   Array of supported video bid response protocols. Refer to List
	//   5.8. At least one supported protocol must be specified in
	//   either the protocol or protocols attribute.
	Protocols []int64 `json:"protocols"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Width of the video player in pixels.
	W int64 `json:"w"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Height of the video player in pixels.
	H int64 `json:"h"`

	// Attribute:
	//   startdelay
	// Type:
	//   integer; recommended
	// Description:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or
	//   post-roll ad placements. Refer to List 5.10 for additional
	//   generic values.
	StartDelay int64 `json:"startdelay"`

	// Attribute:
	//   linearity
	// Type:
	//   integer
	// Description:
	//   Indicates if the impression must be linear, nonlinear, etc. If
	//   none specified, assume all are allowed. Refer to List 5.7.
	Linearity int8 `json:"linearity"`

	// Attribute:
	//   sequence
	// Type:
	//   integer
	// Description:
	//   If multiple ad impressions are offered in the same bid request,
	//   the sequence number will allow for the coordinated delivery
	//   of multiple creatives.
	Sequence int64 `json:"sequence"`

	// Attribute:
	//   sequence
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BAttr []int64 `json:"battr"`

	// Attribute:
	//   maxextended
	// Type:
	//   integer
	// Description:
	//   Maximum extended video ad duration if extension is allowed.
	//   If blank or 0, extension is not allowed. If -1, extension is
	//   allowed, and there is no time limit imposed. If greater than 0,
	//   then the value represents the number of seconds of extended
	//   play supported beyond the maxduration value.
	MaxExtended int64 `json:"maxextended"`

	// Attribute:
	//   minbitrate
	// Type:
	//   integer
	// Description:
	//   Minimum bit rate in Kbps. Exchange may set this dynamically
	//   or universally across their set of publishers.
	MinBitRate int64 `json:"minbitrate"`

	// Attribute:
	//   maxbitrate
	// Type:
	//   integer
	// Description:
	//   Maximum bit rate in Kbps. Exchange may set this dynamically
	//   or universally across their set of publishers.
	MaxBitRate int64 `json:"maxbitrate"`

	// Attribute:
	//   boxingallowed
	// Type:
	//   integer; default 1
	// Description:
	//   Indicates if letter-boxing of 4:3 content into a 16:9 window is
	//   allowed, where 0 = no, 1 = yes.
	BoxingAllowed int8 `json:"boxingallowed"`

	// Attribute:
	//   playbackmethod
	// Type:
	//   integer array
	// Description:
	//   Allowed playback methods. If none specified, assume all are
	//   allowed. Refer to List 5.9.
	PlaybackMethod []int64 `json:"playbackmethod"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Description:
	//   Supported delivery methods (e.g., streaming, progressive). If
	//   none specified, assume all are supported. Refer to List 5.13.
	Delivery []int64 `json:"delivery"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List 5.4
	Pos int64 `json:"pos"`

	// Attribute:
	//   companionad
	// Type:
	//   object array
	// Description:
	//   Array of Banner objects (Section 3.2.3) if companion ads are
	//   available.
	CompanionAd []Banner `json:"companionad"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List 5.6. If an API is not explicitly listed, it is assumed not to be supported.
	API []int64 `json:"api"`

	// Attribute:
	//   companiontype
	// Type:
	//   integer array
	// Description:
	//   Supported VAST companion ad types. Refer to List 5.12.
	//   Recommended if companion Banner objects are included via
	//   the companionad array.
	CompanionType []int64 `json:"companiontype"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}
