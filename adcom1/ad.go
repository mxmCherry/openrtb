package adcom1

import "encoding/json"

// Ad object is the root of a structure that defines in instance of advertising media.
// It includes metadata about the ad overall and sub-objects that provide additional detail specific to the type of media comprising the creative.
type Ad struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Definition:
	//   ID of the creative; unique at least throughout the scope of a vendor (e.g., an exchange or buying platform).
	//   Note that multiple instances of the same ad when used in transactions must have the same ID.
	ID string `json:"id"`

	// Attribute:
	//   adomain
	// Type:
	//   string array; recommended
	// Definition:
	//   Advertiser domain; top two levels only (e.g., “ford.com”).
	//   This can be an array for the case of rotating creatives.
	ADomain []string `json:"adomain,omitempty"`

	// Attribute:
	//   bundle
	// Type:
	//   string array
	// Definition:
	//   When the product of the ad is an app, the unique ID of that app as a bundle or package name (e.g., “com.foo.mygame”).
	//   This should NOT be an app store ID (e.g., no iTunes store IDs).
	//   This can be an array of for the case of rotating creatives.
	Bundle []string `json:"bundle,omitempty"`

	// Attribute:
	//   iurl
	// Type:
	//   string
	// Definition:
	//   URL without cache-busting to an image that is representative of the ad content for cursory level ad quality checking.
	IURL string `json:"iurl,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the ad using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer; default 2
	// Definition:
	//   The taxonomy in use for the cat attribute.
	//   Refer to List: Category Taxonomies.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   lang
	// Type:
	//   string
	// Definition:
	//   Language of the creative using ISO-639-1-alpha-2.
	//   In practice, vendors using this object may elect an alternate standard (e.g., BCP-47) in which case this must be communicated a priori.
	//   The non-standard code “xx” may also be used if the creative has no linguistic content (e.g., a banner with just a company logo).
	Lang string `json:"lang,omitempty"`

	// Attribute:
	//   attr
	// Type:
	//   integer array
	// Definition:
	//   Set of attributes describing the creative.
	//   Refer to List: Creative Attributes.
	Attr []CreativeAttribute `json:"attr,omitempty"`

	// Attribute:
	//   secure
	// Type:
	//   integer
	// Definition:
	//   Flag to indicate if the creative is secure (i.e., uses HTTPS for all assets and markup), where 0 = no, 1 = yes.
	//   There is no default and thus if omitted, the secure state is unknown.
	//   However, as a practical matter, the safe assumption is to treat unknown as non-secure.
	Secure int8 `json:"secure,omitempty"`

	// Attribute:
	//   mrating
	// Type:
	//   integer
	// Definition:
	//   Media rating per IQG guidelines.
	//   Refer to List: Media Ratings.
	MRating MediaRating `json:"mrating,omitempty"`

	// Attribute:
	//   init
	// Type:
	//   integer
	// Definition:
	//   Timestamp of the original instantiation of this ad (i.e., this object or any of its children) in Unix format (i.e., milliseconds since the epoch).
	Init int64 `json:"init,omitempty"`

	// Attribute:
	//   lastmod
	// Type:
	//   integer
	// Definition:
	//   Timestamp of most recent modification to this ad (i.e., this object or any of its children other than the Audit object) in Unix format (i.e., milliseconds since the epoch).
	LastMod int64 `json:"lastmod,omitempty"`

	// Attribute:
	//   display
	// Type:
	//   Object; required *
	// Definition:
	//   Media Subtype Object that indicates this is a display ad and provides additional detail as such.
	//   Refer to Object: Display.
	//   * Required if no other media subtype object is specified.
	Display *Display `json:"display,omitempty"`

	// Attribute:
	//   video
	// Type:
	//   object; required *
	// Definition:
	//   Media Subtype Object that indicates this is a video ad and provides additional detail as such.
	//   Refer to Object: Video.
	//   * Required if no other media subtype object is specified.
	Video *Video `json:"video,omitempty"`

	// Attribute:
	//   audio
	// Type:
	//   object; required *
	// Definition:
	//   Media Subtype Object that indicates this is an audio ad and provides additional detail as such.
	//   Refer to Object: Audio.
	//   * Required if no other media subtype object is specified.
	Audio *Audio `json:"audio,omitempty"`

	// Attribute:
	//   audit
	// Type:
	//   object
	// Definition:
	//   An object depicting the audit status of the ad; typically part of a quality/safety review process.
	//   Refer to Object: Audit.
	Audit *Audit `json:"audit,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
