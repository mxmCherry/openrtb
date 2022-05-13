package adcom1

import "encoding/json"

// BrandVersion provides further identification based on User-Agent Client Hints.
// The BrandVersion object is used to identify a device’s browser or similar software component, and the user agent’s execution platform or operating system.
type BrandVersion struct {
	// Attribute:
	//   brand
	// Type:
	//   string; recommended
	// Definition:
	//   A brand identifier, for example, “Chrome” or “Windows”.
	//   The value may be sourced from the User-Agent Client Hints headers, representing either the user agent brand (from the Sec-CH-UA-Full-Version header) or the platform brand (from the Sec-CH-UA-Platform header).
	Brand string `json:"brand,omitempty"`

	// Attribute:
	//   version
	// Type:
	//   string array
	// Definition:
	//   A sequence of version components, in descending hierarchical order (major, minor, micro, …).
	Version []string `json:"version,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
