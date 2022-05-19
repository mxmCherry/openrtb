package openrtb2

import "encoding/json"

// 3.2.30 Object: BrandVersion
//
// Further identification based on User-Agent Client Hints, the BrandVersion object is used to identify a device’s browser or similar software component, and the user agent’s execution platform or operating system.
type BrandVersion struct {

	// Attribute:
	//   brand
	// Type:
	//   string; required
	// Description:
	//   A brand identifier, for example, "Chrome" or "Windows". The value may be
	//   sourced from the User-Agent Client Hints headers, representing either the
	//   user agent brand (from the Sec-CH-UA-Full-Version header) or the platform
	//   brand (from the Sec-CH-UA-Platform header).
	Brand string `json:"brand"`

	// Attribute:
	//   version
	// Type:
	//   string array
	// Description:
	//   A sequence of version components, in descending hierarchical order (major,
	//   minor, micro, …)
	Version []string `json:"version,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
