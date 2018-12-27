package adcom

import "encoding/json"

// Site object is used to define an ad supported website, in contrast to a non-browser application, for example.
// As a derived class, a Site object inherits all DistributionChannel attributes and adds those defined below.
type Site struct {
	DistributionChannel

	// Attribute:
	//   domain
	// Type:
	//   string
	// Definition:
	//   Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the site using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   sectcat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the current section of the site using IDs from the taxonomy indicated in cattax.
	SectCat []string `json:"sectcat,omitempty"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the current page or view of the site using IDs from the taxonomy indicated in cattax.
	PageCat []string `json:"pagecat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer
	// Definition:
	//   The taxonomy in use for the cat, sectcat and pagecat attributes.
	//   Refer to List: Category Taxonomies.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   privpolicy
	// Type:
	//   integer
	// Definition:
	//   Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivPolicy int8 `json:"privpolicy,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Definition:
	//   Comma separated list of keywords about the site.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   page
	// Type:
	//   string
	// Definition:
	//   URL of the page within the site.
	Page string `json:"page,omitempty"`

	// Attribute:
	//   ref
	// Type:
	//   string
	// Definition:
	//   Referrer URL that caused navigation to the current page.
	Ref string `json:"ref,omitempty"`

	// Attribute:
	//   search
	// Type:
	//   string
	// Definition:
	//   Search string that caused navigation to the current page.
	Search string `json:"search,omitempty"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Definition:
	//   Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes.
	Mobile int8 `json:"mobile,omitempty"`

	// Attribute:
	//   amp
	// Type:
	//   integer
	// Definition:
	//   Indicates if the page is built with AMP HTML, where 0 = no, 1 = yes.
	AMP int8 `json:"amp,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
