package response

import "github.com/mxmCherry/openrtb"

// 5.4 Object: Image
//
// Corresponds to the Image Object in the request.
// The Image object to be used for all image elements of the Native ad such as Icons, Main Image, etc.
type Image struct {
	// Field:
	//   url
	// Scope:
	//   required
	// Type:
	//   string
	// Description:
	//   URL of the image asset
	URL string `json:"url,omitempty"`

	// Field:
	//   w
	// Scope:
	//   recommended
	// Type:
	//   int
	// Description:
	//   Width of the image in pixels
	W int `json:"w,omitempty"`

	// Field:
	//   h
	// Scope:
	//   recommended
	// Type:
	//   int
	// Description:
	//   Height of the image in pixels
	H int `json:"h,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Description:
	//   This object is a placeholder that may contain custom JSON agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext openrtb.RawJSON `json:"ext,omitempty"`
}
