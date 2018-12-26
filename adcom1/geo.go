package adcom

import "encoding/json"

// Geo object encapsulates various methods for specifying a geographic location.
// When subordinate to a Device object, it indicates the location of the device which can also be interpreted as the user's current location.
// When subordinate to a User object, it indicates the location of the user's home base (i.e., not necessarily their current location).
//
// The lat and lon attributes should only be passed if they conform to the accuracy depicted in the type attribute.
// For example, the centroid of a large region (e.g., postal code) should not be passed.
type Geo struct {
	// Attribute:
	//   type
	// Type:
	//   integer
	// Definition:
	//   Source of location data; recommended when passing lat/lon.
	//   Refer to List: Location Types.
	Type interface{} `json:"type,omitempty"`

	// Attribute:
	//   lat
	// Type:
	//   float
	// Definition:
	//   Latitude from -90.0 to +90.0, where negative is south.
	Lat float64 `json:"lat,omitempty"`

	// Attribute:
	//   lon
	// Type:
	//   float
	// Definition:
	//   Longitude from -180.0 to +180.0, where negative is west.
	Lon float64 `json:"lon,omitempty"`

	// Attribute:
	//   accur
	// Type:
	//   integer
	// Definition:
	//   Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device's location services (i.e., type = 1).
	//   Note that this is the accuracy as reported from the device.
	//   Consult OS specific documentation (e.g., Android, iOS) for exact interpretation.
	Accur int `json:"accur,omitempty"`

	// Attribute:
	//   lastfix
	// Type:
	//   integer
	// Definition:
	//   Number of seconds since this geolocation fix was established.
	//   Note that devices may cache location data across multiple fetches.
	//   Ideally, this value should be from the time the actual fix was taken.
	LastFix int64 `json:"lastfix,omitempty"`

	// Attribute:
	//   ipserv
	// Type:
	//   integer
	// Definition:
	//   Service or provider used to determine geolocation from IP address if applicable (i.e., type = 2).
	//   Refer to List: IP Location Services.
	IPServ interface{} `json:"ipserv,omitempty"`

	// Attribute:
	//   country
	// Type:
	//   string
	// Definition:
	//   Country code using ISO-3166-1-alpha-2.
	//   Note that alpha-3 codes may be encountered and vendors are encouraged to be tolerant of them.
	Country string `json:"country,omitempty"`

	// Attribute:
	//   region
	// Type:
	//   string
	// Definition:
	//   Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region,omitempty"`

	// Attribute:
	//   metro
	// Type:
	//   string
	// Definition:
	//   Regional marketing areas such as Nielsen's DMA codes or other similar taxonomy to be agreed among vendors prior to use.
	//   Note that DMA is a trademarked asset of The Nielsen Company.
	//   Vendors are encouraged to ensure their use of DMAs is properly licensed.
	Metro string `json:"metro,omitempty"`

	// Attribute:
	//   city
	// Type:
	//   string
	// Definition:
	//   City using United Nations Code for Trade & Transport Locations “UN/LOCODE” with the space between country and city suppressed (e.g., Boston MA, USA = “USBOS”).
	//   Refer to UN/LOCODE Code List.
	City string `json:"city,omitempty"`

	// Attribute:
	//   zip
	// Type:
	//   string
	// Definition:
	//   ZIP or postal code.
	ZIP string `json:"zip,omitempty"`

	// Attribute:
	//   utcoffset
	// Type:
	//   integer
	// Definition:
	//   Local time as the number +/- of minutes from UTC.
	UTCOffset int `json:"utcoffset,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
