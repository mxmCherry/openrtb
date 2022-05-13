package openrtb2

import "encoding/json"

// 3.2.15 Object: Publisher
//
// This object describes the publisher of the media in which the ad will be displayed.
// The publisher is typically the seller in an OpenRTB transaction.
type Publisher struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Exchange-specific publisher ID.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Publisher name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer; default 1
	// Description:
	//   The taxonomy in use. Refer to the AdCOM list List: Category
	//   Taxonomies for values.
	CatTax int64 `json:"cattax,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the publisher.
	//   The taxonomy to be used is defined by the cattax field. If no
	//   cattax field is supplied IAB Content Category Taxonomy 1.0 is
	//   assumed.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Highest level domain of the publisher (e.g., “publisher.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
