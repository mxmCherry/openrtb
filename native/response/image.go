package response

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
	URL string `json:"url"`

	// Field:
	//   w
	// Scope:
	//   recommended
	// Type:
	//   int
	// Description:
	//   Width of the image in pixels
	W uint64 `json:"w,omitempty"`

	// Field:
	//   h
	// Scope:
	//   recommended
	// Type:
	//   int
	// Description:
	//   Height of the image in pixels
	H uint64 `json:"h,omitempty"`

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
