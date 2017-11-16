// Package response provides OpenRTB Dynamic Native Ads API Specification Version 1.1
// section 5 Native Ad Response Markup Details types:
// https://www.iab.com/guidelines/real-time-bidding-rtb-project/
// https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-Native-Ads-Specification-1-1_2016.pdf
package response

// 5.1 Object: Response
//
// The native object is the top level JSON object which identifies a native response.
type Response struct {
	// Field:
	//   ver
	// Scope:
	//   optional
	// Type:
	//   string
	// Default:
	//   1.1
	// Description:
	//   Version of the Native Markup version in use.
	Ver string `json:"ver,omitempty"`

	// Field:
	//   assets
	// Scope:
	//   required
	// Type:
	//   object array
	// Description:
	//   List of native ad’s assets.
	Assets []Asset `json:"assets"`

	// Field:
	//   link
	// Scope:
	//   required
	// Type:
	//   object
	// Description:
	//   Destination Link. This is default link object for the ad.
	//   Individual assets can also have a link object which applies if the asset is activated(clicked).
	//   If the asset doesn’t have a link object, the parent link object applies.
	Link Link `json:"link"`

	// Field:
	//   imptrackers
	// Scope:
	//   optional
	// Type:
	//   string array
	// Description:
	//   Array of impression tracking URLs, expected to return a 1x1 image or 204 response - typically
	//   only passed when using 3rd party trackers.
	ImpTrackers []string `json:"imptrackers,omitempty"`

	// Field:
	//   jstracker
	// Scope:
	//   optional
	// Type:
	//   string
	// Description:
	//   Optional JavaScript impression tracker. This is a valid HTML, Javascript is already wrapped in <script> tags.
	//   It should be executed at impression time where it can be supported.
	JSTracker string `json:"jstracker,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	//   This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext RawJSON `json:"ext,omitempty"`
}
