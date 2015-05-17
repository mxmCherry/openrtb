package rtb

// 5.7 Video Linearity
//
// The following table indicates the options for video linearity. “In-stream” or “linear” video refers to preroll,
// post-roll, or mid-roll video ads where the user is forced to watch ad in order to see the video
// content. “Overlay” or “non-linear” refer to ads that are shown on top of the video content.
//
// This field is optional. The following is the interpretation of the bidder based upon the presence or
// absence of the field in the bid request:
// - If no value is set, any ad (linear or not) can be present in the response.
// - If a value is set, only ads of the corresponding type can be present in the response.
//
// Note to the reader: This OpenRTB table has values derived from the IAB Quality Assurance Guidelines
// (QAG). Practitioners should keep in sync with updates to the QAG values as published on IAB.net.
const (
	VideoLinearityLinear    = 1 // 1 Linear / In-Stream
	VideoLinearityNonLinear = 2 // 2 Non-Linear / Overlay
)
