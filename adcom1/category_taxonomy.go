package adcom

// CategoryTaxonomy identifies the taxonomy in effect when content categories are listed.
type CategoryTaxonomy int

// CategoryTaxonomy options.
//
// Values of 500+ hold vendor-specific codes.
const (
	CatTaxIABContent1 CategoryTaxonomy = 1 // 1	IAB Content Category Taxonomy 1.0.
	CatTaxIABContent2 CategoryTaxonomy = 2 // 2	IAB Content Category Taxonomy 2.0: www.iab.com/guidelines/taxonomy
	CatTaxIABProduct1 CategoryTaxonomy = 3 // 3	IAB Ad Product Taxonomy 1.0.
)
