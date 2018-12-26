package adcom

import "encoding/json"

// Restrictions object allows lists of restrictions on ad responses to be specified including specific content categories, advertisers, ads pertaining to specific apps, or creative attributes.
type Restrictions struct {
	// Attribute:
	//   bcat
	// Type:
	//   string array
	// Definition:
	//   Block list of content categories using IDs from the taxonomy indicated in cattax.
	BCat []string `json:"bcat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer; default 2
	// Definition:
	//   The taxonomy in use for the bcat attribute
	//   Refer to List: Category Taxonomies.
	CatTax interface{} `json:"cattax,omitempty"`

	// Attribute:
	//   badv
	// Type:
	//   string array
	// Definition:
	//   Block list of advertisers by their domains (e.g., “ford.com”).
	BAdv []string `json:"badv,omitempty"`

	// Attribute:
	//   bapp
	// Type:
	//   string array
	// Definition:
	//   Block list of apps for which ads are disallowed
	//   These should be bundle or package names (e.g., “com.foo.mygame”) and should NOT be app store IDs (e.g., not iTunes store IDs).
	BApp []string `json:"bapp,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Definition:
	//   Block list of creative attributes
	//   Refer to List: Creative Attributes.
	BAttr []interface{} `json:"battr,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
