// rtb package provides common OpenRTB's Object types and constants,
// according to OpenRTB API Specification Version 2.3:
// http://openrtb.github.io/OpenRTB/
package rtb


// Gender, where “M” = male, “F” = female, “O” = known to be
//   other (i.e., omitted is unknown).
const (
	UserGenderMale    string = "M" // “M” = male
	UserGenderFemale  string = "F" // “F” = female
	UserGenderOther   string = "O" // “O” = known to be other
	UserGenderUnknown string = ""  // omitted is unknown
)


// 5.16 Location Type
//   The following table lists the options to indicate how the geographic information was determined.
const (
	GeoTypeGPS          uint8 = 1 // 1 GPS/Location Services
	GeoTypeIP           uint8 = 2 // 2 IP Address
	GeoTypeUserProvided uint8 = 3 // 3 User provided (e.g., registration data)
)


// Mobile-optimized signal, where 0 = no, 1 = yes.
const (
	SiteMobileNo  uint8 = 0 // 0 = no
	SiteMobileYes uint8 = 1 // 1 = yes
)


// Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
const (
	SitePrivacyPolicyNo  uint8 = 0 // 0 = no
	SitePrivacyPolicyYes uint8 = 1 // 1 = yes
)


// Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
const (
	AppPrivacyPolicyNo  uint8 = 0 // 0 = no
	AppPrivacyPolicyYes uint8 = 1 // 1 = yes
)


// 0 = not live, 1 = content is live (e.g., stream, live blog).
const (
	ContentLiveStreamNo  uint8 = 0 // 0 = not live
	ContentLiveStreamYes uint8 = 1 // 1 = content is live (e.g., stream, live blog)
)


// 0 = indirect, 1 = direct.
const (
	ContentSourceRelationshipIndirect uint8 = 0 // 0 = indirect
	ContentSourceRelationshipDirect   uint8 = 1 // 1 = direct
)

// Indicator of whether or not the content is embeddable (e.g.,
// an embeddable video player), where 0 = no, 1 = yes.
const (
	ContentEmbeddableNo  uint8 = 0 // 0 = no
	ContentEmbeddableYes uint8 = 1 // 1 = yes
)


// Placeholder for exchange-specific extensions to OpenRTB.
type Ext interface {}


// 3.2.6 Object: Site
//   This object should be included if the ad supported content is a website as opposed to a non-browser
//   application. A bid request must not contain both a Site and an App object. At a minimum, it is useful
//   to provide a site ID or page URL, but this is not strictly required.
type Site struct {
	
	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific site ID.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Site name (may be aliased at the publisher’s request).
	Name string `json:"name"`
	
	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain"`
	
	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories of the site. Refer to List 5.1.
	Cat []string `json:"cat"`
	
	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current
	//   section of the site. Refer to List 5.1.
	SectionCat []string `json:"sectioncat"`
	
	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current page
	//   or view of the site. Refer to List 5.1.
	PageCat []string `json:"pagecat"`
	
	// Attribute:
	//   page
	// Type:
	//   string
	// Description:
	//   URL of the page where the impression will be shown.
	Page string `json:"page"`
	
	// Attribute:
	//   ref
	// Type:
	//   string
	// Description:
	//   Referrer URL that caused navigation to the current page.
	Ref string `json:"ref"`
	
	// Attribute:
	//   search
	// Type:
	//   string
	// Description:
	//   Search string that caused navigation to the current page.
	Search string `json:"search"`
	
	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Description:
	//   Mobile-optimized signal, where 0 = no, 1 = yes.
	Mobile uint8 `json:"mobile"`
	
	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy uint8 `json:"privacypolicy"`
	
	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.8) of the site.
	Publisher Publisher `json:"publisher"`
	
	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.9) within the site.
	Content Content `json:"content"`
	
	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the site.
	Keywords string `json:"keywords"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.7 Object: App
//   This object should be included if the ad supported content is a non-browser application (typically in
//   mobile) as opposed to a website. A bid request must not contain both an App and a Site object. At a
//   minimum, it is useful to provide an App ID or bundle, but this is not strictly required.
type App struct {
	
	// Attrubute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific app ID.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Tyoe:
	//   string
	// Description:
	//   App name (may be aliased at the publisher’s request).
	Name string `json:"name"`
	
	// Attribute:
	//   bundle
	// Type:
	//   string
	// Description:
	//   Application bundle or package name (e.g., com.foo.mygame);
	//   intended to be a unique ID across exchanges.
	Bundle string `json:"bundle"`
	
	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain"`
	
	// Attribute:
	//   storeurl
	// Type:
	//   string
	// Description:
	//   App store URL for an installed app; for QAG 1.5 compliance.
	StoreURL string `json:"storeurl"`
	
	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories of the app. Refer to List 5.1.
	Cat []string `json:"cat"`
	
	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current
	//   section of the app. Refer to List 5.1.
	SectionCat []string `json:"sectioncat"`
	
	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current page
	//   or view of the app. Refer to List 5.1.
	PageCat []string `json:"pagecat"`
	
	// Attribute:
	//   ver
	// Type:
	//   string
	// Description:
	//   Application version.
	Ver string `json:"ver"`
	
	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy uint8 `json:"privacypolicy"`
	
	// Attribute:
	//   paid
	// Type:
	//   integer
	// Description:
	//   0 = app is free, 1 = the app is a paid version.
	Paid uint8 `json:"paid"`
	
	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.8) of the app.
	Publisher Publisher `json:"publisher"`
	
	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.9) within the app.
	Content Content `json:"content"`
	
	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the app.
	Keywords string `json:"keywords"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.8 Object: Publisher
//   This object describes the publisher of the media in which the ad will be displayed. The publisher is
//   typically the seller in an OpenRTB transaction.
type Publisher struct {
	
	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Exchange-specific publisher ID.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Publisher name (may be aliased at the publisher’s request).
	Name string `json:"name"`
	
	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the publisher.
	//   Refer to List 5.1.
	Cat []string `json:"cat"`
	
	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Highest level domain of the publisher (e.g., “publisher.com”).
	Domain string `json:"domain"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.9 Object: Content
//   This object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
//   content. This object may be useful when syndicated content contains impressions and does
//   not necessarily match the publisher’s general content. The exchange might or might not have
//   knowledge of the page where the content is running, as a result of the syndication method. For
//   example might be a video impression embedded in an iframe on an unknown web property or device.
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
	Producer Producer `json:"producer"`
	
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
	VideoQuality uint64 `json:"videoquality"`
	
	// Attribute:
	//   context
	// Type:
	//   integer
	// Description:
	//   Type of content (game, video, text, etc.). Refer to List 5.14.
	Context uint64 `json:"context"`
	
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
	QAGMediaRating uint64 `json:"qagmediarating"`
	
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


// 3.2.10 Object: Producer
//   This object defines the producer of the content in which the ad will be shown. This is particularly useful
//   when the content is syndicated and may be distributed through different publishers and thus when the
//   producer and publisher are not necessarily the same entity.
type Producer struct {
	
	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Content producer or originator ID. Useful if content is
	//   syndicated and may be posted on a site using embed tags.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Content producer or originator name (e.g., “Warner Bros”).
	Name string `json:"name"`
	
	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the content
	//   producer. Refer to List 5.1.
	Cat []string `json:"cat"`
	
	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Highest level domain of the content producer (e.g.,
	//   “producer.com”).
	Domain string `json:"domain"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.12 Object: Geo
//   This object encapsulates various methods for specifying a geographic location. When subordinate to a
//   Device object, it indicates the location of the device which can also be interpreted as the user’s current
//   location. When subordinate to a User object, it indicates the location of the user’s home base (i.e., not
//   necessarily their current location).
//   The lat/lon attributes should only be passed if they conform to the accuracy depicted in the type
//   attribute. For example, the centroid of a geographic region such as postal code should not be passed.
type Geo struct {
	
	// Attribute:
	//   lat
	// Type:
	//   float
	// Description:
	//   Latitude from -90.0 to +90.0, where negative is south.
	Lat float64 `json:"lat"`
	
	// Attribute:
	//   lon
	// Type:
	//   float
	// Description:
	//   Longitude from -180.0 to +180.0, where negative is west.
	Lon float64 `json:"lon"`
	
	// Attribute:
	//   type
	// Type:
	//   integer
	// Description:
	//   Source of location data; recommended when passing
	//   lat/lon. Refer to List 5.16.
	Type uint8 `json:"type"`
	
	// Attribute:
	//   country
	// Type:
	//   string
	// Description:
	//   Country code using ISO-3166-1-alpha-3.
	Country string `json:"country"`
	
	// Attribute:
	//   region
	// Type:
	//   string
	// Description:
	//   Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region"`
	
	// Attribute:
	//   regionfips104
	// Type:
	//   string
	// Description:
	//   Region of a country using FIPS 10-4 notation. While OpenRTB
	//   supports this attribute, it has been withdrawn by NIST in 2008.
	RegionFIPS104 string `json:"regionfips104"`
	
	// Attribute:
	//   metro
	// Type:
	//   string
	// Description:
	//   Google metro code; similar to but not exactly Nielsen DMAs.
	//   See Appendix A for a link to the codes.
	Metro string `json:"metro"`
	
	// Attribute:
	//   city
	// Type:
	//   string
	// Description:
	//   City using United Nations Code for Trade & Transport
	//   Locations. See Appendix A for a link to the codes.
	City string `json:"city"`
	
	// Attribute:
	//   zip
	// Type:
	//   string
	// Description:
	//   Zip or postal code.
	ZIP string `json:"zip"`
	
	// Attribute:
	//   utcoffset
	// Type:
	//   integer
	// Description:
	//   Local time as the number +/- of minutes from UTC.
	UTCOffset int8 `json:"utcoffset"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.13 Object: User
//   This object contains information known or derived about the human user of the device (i.e., the
//   audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
//   privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
//   frequency capping and retargeting.
type User struct {
	
	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific ID for the user. At least one of id or
	//   buyerid is recommended.
	ID string `json:"id"`
	
	// Attribute:
	//   buyerid
	// Type:
	//   string; recommended
	// Description:
	//   Buyer-specific ID for the user as mapped by the exchange for
	//   the buyer. At least one of buyerid or id is recommended.
	BuyerID string `json:"buyerid"`
	
	// Attribute:
	//   yob
	// Type:
	//   integer
	// Description:
	//   Year of birth as a 4-digit integer.
	Yob uint16 `json:"yob`
	
	// Attribute:
	//   gender
	// Type:
	//   string
	// Description:
	//   Gender, where “M” = male, “F” = female, “O” = known to be
	//   other (i.e., omitted is unknown).
	Gender string `json:"gender"`
	
	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords, interests, or intent.
	Keywords string `json:"keywords"`
	
	// Attribute:
	//   customdata
	// Type:
	//   string
	// Description:
	//   Optional feature to pass bidder data that was set in the
	//   exchange’s cookie. The string must be in base85 cookie safe
	//   characters and be in any format. Proper JSON encoding must
	//   be used to include “escaped” quotation marks.
	CustomData string `json:"customdata"`
	
	// Attribute:
	//   geo
	// Type
	//   object
	// Description:
	//   Location of the user’s home base defined by a Geo object
	//   (Section 3.2.12). This is not necessarily their current location.
	Geo Geo `json:"geo"`
	
	// Attribute:
	//   data
	// Type:
	//   object array
	// Description:
	//   Additional user data. Each Data object (Section 3.2.14)
	//   represents a different data source.
	Data Data `json:"data"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.14 Object: Data
//   The data and segment objects together allow additional data about the user to be specified. This data
//   may be from multiple sources whether from the exchange itself or third party providers as specified by
//   the id field. A bid request can mix data objects from multiple providers. The specific data providers in
//   use should be published by the exchange a priori to its bidders.
type Data struct {
	
	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Exchange-specific ID for the data provider.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Exchange-specific name for the data provider.
	Name string `json:"name"`
	
	// Attribute:
	//   segment
	// Type:
	//   object array
	// Description:
	//   Array of Segment (Section 3.2.15) objects that contain the
	//   actual data values.
	Segment []Segment `json:"segment"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}


// 3.2.15 Object: Segment
//   Segment objects are essentially key-value pairs that convey specific units of data about the user. The
//   parent Data object is a collection of such values from a given data provider. The specific segment
//   names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	
	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   ID of the data segment specific to the data provider.
	ID string `json:"id"`
	
	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Name of the data segment specific to the data provider.
	Name string `json:"name"`
	
	// Attribute:
	//   value
	// Type:
	//   string
	// Description:
	//   String representation of the data segment value.
	Value string `json:"value"`
	
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
	
}