package adcom1

import "encoding/json"

// Content object describes the content in which an impression can appear, which may be syndicated or non-syndicated content.
// This object may be useful when syndicated content contains impressions and does not necessarily match the publisher's general content.
// An exchange may or may not have knowledge of the page where the content is running as a result of the syndication method (e.g., a video impression embedded in an iframe on an unknown web property or device).
type Content struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   ID uniquely identifying the content.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   episode
	// Type:
	//   integer
	// Definition:
	//   Episode number.
	Episode int64 `json:"episode,omitempty"`

	// Attribute:
	//   title
	// Type:
	//   string
	// Definition:
	//   Content title.
	//   Video Examples: “Search Committee” (television), “Star Wars, A New Hope” (movie), or “Endgame” (made for web).
	//   Non-Video Example: “Why an Antarctic Glacier Is Melting So Quickly” (Time magazine article).
	Title string `json:"title,omitempty"`

	// Attribute:
	//   series
	// Type:
	//   string
	// Definition:
	//   Content series.
	//   Video Examples: “The Office” (television), “Star Wars” (movie), or “Arby 'N' The Chief” (made for web).
	//   Non-Video Example: “Ecocentric” (Time Magazine blog).
	Series string `json:"series,omitempty"`

	// Attribute:
	//   season
	// Type:
	//   string
	// Definition:
	//   Content season (e.g., “Season 3”).
	Season string `json:"season,omitempty"`

	// Attribute:
	//   artist
	// Type:
	//   string
	// Definition:
	//   Artist credited with the content.
	Artist string `json:"artist,omitempty"`

	// Attribute:
	//   genre
	// Type:
	//   string
	// Definition:
	//   Genre that best describes the content (e.g., rock, pop, etc).
	Genre string `json:"genre,omitempty"`

	// Attribute:
	//   album
	// Type:
	//   string
	// Definition:
	//   Album to which the content belongs; typically for audio.
	Album string `json:"album,omitempty"`

	// Attribute:
	//   isrc
	// Type:
	//   string
	// Definition:
	//   International Standard Recording Code conforming to ISO-3901.
	ISRC string `json:"isrc,omitempty"`

	// Attribute:
	//   url
	// Type:
	//   string
	// Definition:
	//   URL of the content, for buy-side contextualization or review.
	URL string `json:"url,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Definition:
	//   Array of content categories describing the content using IDs from the taxonomy indicated in cattax.
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
	//   prodq
	// Type:
	//   integer
	// Definition:
	//   Production quality.
	//   Refer to List: Production Qualities.
	ProdQ ProductionQuality `json:"prodq,omitempty"`

	// Attribute:
	//   context
	// Type:
	//   integer
	// Definition:
	//   Type of content (game, video, text, etc.).
	//   Refer to List: Content Contexts.
	Context ContentContext `json:"context,omitempty"`

	// Attribute:
	//   rating
	// Type:
	//   string
	// Definition:
	//   Content rating (e.g., MPAA).
	Rating string `json:"rating,omitempty"`

	// Attribute:
	//   urating
	// Type:
	//   string
	// Definition:
	//   User rating of the content (e.g., number of stars, likes, etc.).
	URating string `json:"urating,omitempty"`

	// Attribute:
	//   mrating
	// Type:
	//   integer
	// Definition:
	//   Media rating per IQG guidelines.
	//   Refer to List: Media Ratings.
	MRating MediaRating `json:"mrating,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Definition:
	//   Comma separated list of keywords describing the content.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   kwarray
	// Type:
	//   string array
	// Definition:
	//   Array of keywords about the site. Only one of 'keywords' or 'kwarray' may be present.
	KwArray []string `json:"kwarray,omitempty"`

	// Attribute:
	//   live
	// Type:
	//   integer
	// Definition:
	//   Indication of live content, where 0 = not live, 1 = live (e.g., stream, live blog).
	Live int8 `json:"live,omitempty"`

	// Attribute:
	//   srcrel
	// Type:
	//   integer
	// Definition:
	//   Source relationship, where 0 = indirect, 1 = direct.
	SrcRel int8 `json:"srcrel,omitempty"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Definition:
	//   Length of content in seconds; typically for video or audio.
	Len int64 `json:"len,omitempty"`

	// Attribute:
	//   lang
	// Type:
	//   string
	// Definition:
	//   Content language using ISO-639-1-alpha-2.
	Lang string `json:"lang,omitempty"`

	// Attribute:
	//   embed
	// Type:
	//   integer
	// Definition:
	//   Indicator of whether or not the content is embedded off-site from the the site or app described in those objects (e.g., an embedded video player), where 0 = no, 1 = yes.
	Embed int8 `json:"embed,omitempty"`

	// Attribute:
	//   producer
	// Type:
	//   object
	// Definition:
	//   Details about the content producer.
	//   Refer to Object: Producer.
	Producer *Producer `json:"producer,omitempty"`

	// Attribute:
	//   network
	// Type:
	//   object
	// Definition:
	//   Details about the network.
	//   Refer to Object: Network.
	Network *Network `json:"network,omitempty"`

	// Attribute:
	//   channel
	// Type:
	//   object
	// Definition:
	//   Details about the channel.
	//   Refer to Object: Channel.
	Channel *Channel `json:"channel,omitempty"`

	// Attribute:
	//   data
	// Type:
	//   object array
	// Definition:
	//   Additional user data.
	//   Each Data object represents a different data source.
	//   Refer to Object: Data.
	Data []Data `json:"data,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
