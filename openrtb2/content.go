package openrtb2

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/v16/adcom1"
)

// 3.2.16 Object: Content
//
// This object describes the content in which the impression will appear, which may be syndicated or nonsyndicated content.
// This object may be useful when syndicated content contains impressions and does not necessarily match the publisher’s general content.
// The exchange might or might not have knowledge of the page where the content is running, as a result of the syndication method.
// For example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   ID uniquely identifying the content.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   episode
	// Type:
	//   integer
	// Description:
	//   Episode number.
	Episode int64 `json:"episode,omitempty"`

	// Attribute:
	//   title
	// Type:
	//   string
	// Description:
	//   Content title.
	//   Video Examples: “Search Committee” (television), “A New
	//   Hope” (movie), or “Endgame” (made for web).
	//   Non-Video Example: “Why an Antarctic Glacier Is Melting So
	//   Quickly” (Time magazine article).
	Title string `json:"title,omitempty"`

	// Attribute:
	//   series
	// Type:
	//   string
	// Description:
	//   Content series.
	//   Video Examples: “The Office” (television), “Star Wars” (movie),
	//   or “Arby ‘N’ The Chief” (made for web).
	//   Non-Video Example: “Ecocentric” (Time Magazine blog).
	Series string `json:"series,omitempty"`

	// Attribute:
	//   season
	// Type:
	//   string
	// Description:
	//   Content season (e.g., “Season 3”).
	Season string `json:"season,omitempty"`

	// Attribute:
	//   artist
	// Type:
	//   string
	// Description:
	//   Artist credited with the content.
	Artist string `json:"artist,omitempty"`

	// Attribute:
	//   genre
	// Type:
	//   string
	// Description:
	//   Genre that best describes the content (e.g., rock, pop, etc).
	Genre string `json:"genre,omitempty"`

	// Attribute:
	//   album
	// Type:
	//   string
	// Description:
	//   Album to which the content belongs; typically for audio.
	Album string `json:"album,omitempty"`

	// Attribute:
	//   isrc
	// Type:
	//   string
	// Description:
	//   International Standard Recording Code conforming to ISO-
	//   3901.
	ISRC string `json:"isrc,omitempty"`

	// Attribute:
	//   producer
	// Type:
	//   object
	// Description:
	//   Details about the content Producer (Section 3.2.17).
	Producer *Producer `json:"producer,omitempty"`

	// Attribute:
	//   url
	// Type:
	//   string
	// Description:
	//   URL of the content, for buy-side contextualization or review.
	URL string `json:"url,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer
	// Description:
	//   The taxonomy in use. Refer to the AdCOM list List: Category
	//   Taxonomies for values.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the content.
	//   The taxonomy to be used is defined by the cattax field. If no
	//   cattax field is supplied IAB Content Category Taxonomy 1.0 is
	//   assumed.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   prodq
	// Type:
	//   integer
	// Description:
	//   Production quality. Refer to List: Production Qualities in AdCOM 1.0.
	ProdQ adcom1.ProductionQuality `json:"prodq,omitempty"`

	// Attribute:
	//   context
	// Type:
	//   integer
	// Description:
	//   Type of content (game, video, text, etc.). Refer to List: Content
	//   Contexts in AdCOM 1.0.
	Context adcom1.ContentContext `json:"context,omitempty"`

	// Attribute:
	//   contentrating
	// Type:
	//   string
	// Description:
	//   Content rating (e.g., MPAA).
	ContentRating string `json:"contentrating,omitempty"`

	// Attribute:
	//   userrating
	// Type:
	//   string
	// Description:
	//   User rating of the content (e.g., number of stars, likes, etc.).
	UserRating string `json:"userrating,omitempty"`

	// Attribute:
	//   qagmediarating
	// Type:
	//   integer
	// Description:
	//   Media rating per IQG guidelines. Refer to List: Media Ratings in
	//   AdCOM 1.0.
	QAGMediaRating adcom1.MediaRating `json:"qagmediarating,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords describing the content. Only
	//   one of ‘keywords’ or ‘kwarray’ may be present.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   kwarray
	// Type:
	//   string
	// Description:
	//   Array of keywords about the site. Only one of ‘keywords’ or
	//   ‘kwarray’ may be present.
	KwArray []string `json:"kwarray,omitempty"`

	// Attribute:
	//   livestream
	// Type:
	//   integer
	// Description:
	//   0 = not live, 1 = content is live (e.g., stream, live blog).
	LiveStream int8 `json:"livestream,omitempty"`

	// Attribute:
	//   sourcerelationship
	// Type:
	//   integer
	// Description:
	//   0 = indirect, 1 = direct.
	SourceRelationship int8 `json:"sourcerelationship,omitempty"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Description:
	//   Length of content in seconds; appropriate for video or audio.
	Len int64 `json:"len,omitempty"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Content language using ISO-639-1-alpha-2. Only one of
	//   language or langb should be present.
	Language string `json:"language,omitempty"`

	// Attribute:
	//   langb
	// Type:
	//   string
	// Description:
	//   Content language using IETF BCP 47. Only one of language or
	//   langb should be present.
	LangB string `json:"langb,omitempty"`

	// Attribute:
	//   embeddable
	// Type:
	//   integer
	// Description:
	//   Indicator of whether or not the content is embeddable (e.g.,
	//   an embeddable video player), where 0 = no, 1 = yes.
	Embeddable int8 `json:"embeddable,omitempty"`

	// Attribute:
	//   data
	// Type:
	//   object array
	// Description:
	//   Additional content data. Each Data object (Section 3.2.21)
	//   represents a different data source.
	Data []Data `json:"data,omitempty"`

	// Attribute:
	//   network
	// Type:
	//   object
	// Description:
	//   Details about the network (Section 3.2.23) the content is on.
	Network *Network `json:"network,omitempty"`

	// Attribute:
	//   channel
	// Type:
	//   object
	// Description:
	//   Details about the channel (Section 3.2.24) the content is on.
	Channel *Channel `json:"channel,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
