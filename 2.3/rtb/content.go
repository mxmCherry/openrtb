package rtb

// 3.2.9 Object: Content
//
// This object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher’s general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   ID uniquely identifying the content.
	ID string `json:"id"`

	// Attribute:
	//   episode
	// Type:
	//   integer
	// Description:
	//   Episode number (typically applies to video content).
	Episode uint64 `json:"episode"`

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
	Title string `json:"title"`

	// Attribute:
	//   series
	// Type:
	//   string
	// Description:
	//   Content series.
	//   Video Examples: “The Office” (television), “Star Wars” (movie),
	//   or “Arby ‘N’ The Chief” (made for web).
	//   Non-Video Example: “Ecocentric” (Time Magazine blog).
	Series string `json:"series"`

	// Attribute:
	//   season
	// Type:
	//   string
	// Description:
	//   Content season; typically for video content (e.g., “Season 3”).
	Season string `json:"season"`

	// Attribute:
	//   producer
	// Type:
	//   object
	// Description:
	//   Details about the content Producer (Section 3.2.10).
	Producer *Producer `json:"producer"`

	// Attribute:
	//   url
	// Type:
	//   string
	// Description:
	//   URL of the content, for buy-side contextualization or review.
	URL string `json:"url"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the content
	//   producer. Refer to List 5.1.
	Cat []string `json:"cat"`

	// Attribute:
	//   videoquality
	// Type:
	//   integer
	// Description:
	//   Video quality per IAB’s classification. Refer to List 5.11.
	VideoQuality uint8 `json:"videoquality"`

	// Attribute:
	//   context
	// Type:
	//   integer
	// Description:
	//   Type of content (game, video, text, etc.). Refer to List 5.14.
	Context uint8 `json:"context"`

	// Attribute:
	//   contentrating
	// Type:
	//   string
	// Description:
	//   Content rating (e.g., MPAA).
	ContentRating string `json:"contentrating"`

	// Attribute:
	//   userrating
	// Type:
	//   string
	// Description:
	//   User rating of the content (e.g., number of stars, likes, etc.).
	UserRating string `json:"userrating"`

	// Attribute:
	//   qagmediarating
	// Type:
	//   integer
	// Description:
	//   Media rating per QAG guidelines. Refer to List 5.15.
	QAGMediaRating uint8 `json:"qagmediarating"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords describing the content.
	Keywords string `json:"keywords"`

	// Attribute:
	//   livestream
	// Type:
	//   integer
	// Description:
	//   0 = not live, 1 = content is live (e.g., stream, live blog).
	LiveStream uint8 `json:"livestream"`

	// Attribute:
	//   sourcerelationship
	// Type:
	//   integer
	// Description:
	//   0 = indirect, 1 = direct.
	SourceRelationship uint8 `json:"sourcerelationship"`

	// Attribute:
	//   len
	// Type:
	//   integer
	// Description:
	//   Length of content in seconds; appropriate for video or audio.
	Len uint64 `json:"len"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Content language using ISO-639-1-alpha-2.
	Language string `json:"language"`

	// Attribute:
	//   embeddable
	// Type:
	//   integer
	// Description:
	//   Indicator of whether or not the content is embeddable (e.g.,
	//   an embeddable video player), where 0 = no, 1 = yes.
	Embeddable uint8 `json:"embeddable"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}
