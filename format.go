package openrtb

// 3.2.10 Object: Format
//
// OpenRTB 2.5:
// This object represents an allowed size (i.e., height and width combination) or Flex Ad parameters for a banner impression.
// These are typically used in an array where multiple sizes are permitted. It is recommended that either the w/h pair or
// the wratio/hratio/wmin set (i.e., for Flex Ads) be specified.
type Format struct {
	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Width in device independent pixels (DIPS).
    W uint64 `json:"w"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Height in device independent pixels (DIPS).
    H uint64 `json:"h"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Description:
	//   Relative width when expressing size as a ratio.
    WRatio int `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   integer
	// Description:
	//   Relative height when expressing size as a ratio.
    HRatio int `json:"hratio,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer
	// Description:
	//   The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
    WMin int `json:"wmin,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext RawJSON `json:"ext,omitempty"`
}
