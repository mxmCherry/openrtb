// rtb package provides common OpenRTB's Object types and constants,
// according to OpenRTB API Specification Version 2.3:
// http://openrtb.github.io/OpenRTB/
package rtb


// Gender, where “M” = male, “F” = female, “O” = known to be
//   other (i.e., omitted is unknown).
const (
	UserGenderMale    string = "M"
	UserGenderFemale  string = "F"
	UserGenderOther   string = "O"
	UserGenderUnknown string = ""
)


// 5.16 Location Type
//   The following table lists the options to indicate how the geographic information was determined.
const (
	GeoTypeGPS          uint8 = 1 // 1 GPS/Location Services
	GeoTypeIP           uint8 = 2 // 2 IP Address
	GeoTypeUserProvided uint8 = 3 // 3 User provided (e.g., registration data)
)


// Placeholder for exchange-specific extensions to OpenRTB.
type Ext interface {}


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