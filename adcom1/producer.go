package adcom1

import "encoding/json"

// Producer object defines the producer of the content in which ad will be displayed.
// This is particularly useful when the content is syndicated and may be distributed through different publishers and thus when the producer and publisher are not necessarily the same entity.
type Producer struct {
	// Attribute:
	//   id
	// Type:
	//   string, recommended
	// Definition:
	//   Vendor-specific unique producer identifier.
	//   Useful if content is syndicated and may be posted on a site using embed tags.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Definition:
	//   Displayable name of the producer.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Definition:
	//   Highest level domain of the producer (e.g., “producer.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories that describe the producer using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer
	// Definition:
	//   The taxonomy in use for the cat attribute.
	//   Refer to List: Category Taxonomies.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
