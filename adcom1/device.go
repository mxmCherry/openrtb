package adcom1

import "encoding/json"

// Device object provides information pertaining to the device through which the user is interacting.
// Device information includes its hardware, platform, location, and carrier data.
// The device can refer to a mobile handset, a desktop computer, set top box, or other digital device.
type Device struct {
	// Attribute:
	//   type
	// Type:
	//   integer
	// Definition:
	//   The general type of device.
	//   Refer to List: Device Types.
	Type DeviceType `json:"type,omitempty"`

	// Attribute:
	//   ua
	// Type:
	//   string
	// Definition:
	//   Browser user agent string.
	UA string `json:"ua,omitempty"`

	// Attribute:
	//   ifa
	// Type:
	//   string
	// Definition:
	//   ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`

	// Attribute:
	//   dnt
	// Type:
	//   integer
	// Definition:
	//   Standard “Do Not Track” flag as set in the header by the browser, where 0 = tracking is unrestricted, 1 = do not track.
	DNT int8 `json:"dnt,omitempty"`

	// Attribute:
	//   lmt
	// Type:
	//   integer
	// Definition:
	//   “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, Android), where 0 = tracking is unrestricted, 1 = tracking must be limited per commercial guidelines.
	Lmt int8 `json:"lmt,omitempty"`

	// Attribute:
	//   make
	// Type:
	//   string
	// Definition:
	//   Device make (e.g., "Apple").
	Make string `json:"make,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Definition:
	//   Device model (e.g., “iPhone10,1” when the specific device model is known, “iPhone” otherwise).
	//   The value obtained from the device O/S should be used when available.
	Model string `json:"model,omitempty"`

	// Attribute:
	//   os
	// Type:
	//   integer
	// Definition:
	//   Device operating system.
	//   Refer to List: Operating Systems.
	OS OperatingSystem `json:"os,omitempty"`

	// Attribute:
	//   osv
	// Type:
	//   string
	// Definition:
	//   Device operating system version (e.g., “3.1.2”).
	OSV string `json:"osv,omitempty"`

	// Attribute:
	//   hwv
	// Type:
	//   string
	// Definition:
	//   Hardware version of the device (e.g., “5S” for iPhone 5S).
	HWV string `json:"hwv,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Physical height of the screen in pixels.
	H uint64 `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Physical width of the screen in pixels.
	W uint64 `json:"w,omitempty"`

	// Attribute:
	//   ppi
	// Type:
	//   integer
	// Definition:
	//   Screen size as pixels per linear inch.
	PPI int `json:"ppi,omitempty"`

	// Attribute:
	//   pxratio
	// Type:
	//   float
	// Definition:
	//   The ratio of physical pixels to device independent pixels.
	PxRatio float64 `json:"pxratio,omitempty"`

	// Attribute:
	//   js
	// Type:
	//   integer
	// Definition:
	//   Support for JavaScript, where 0 = no, 1 = yes.
	JS int8 `json:"js,omitempty"`

	// Attribute:
	//   lang
	// Type:
	//   string
	// Definition:
	//   Browser language using ISO-639-1-alpha-2.
	Lang string `json:"lang,omitempty"`

	// Attribute:
	//   ip
	// Type:
	//   string
	// Definition:
	//   IPv4 address closest to device.
	IP string `json:"ip,omitempty"`

	// Attribute:
	//   ipv6
	// Type:
	//   string
	// Definition:
	//   IP address closest to device as IPv6.
	IPv6 string `json:"ipv6,omitempty"`

	// Attribute:
	//   xff
	// Type:
	//   string
	// Definition:
	//   The value of the “x-forwarded-for” header.
	XFF string `json:"xff,omitempty"`

	// Attribute:
	//   iptr
	// Type:
	//   integer
	// Definition:
	//   Indicator of truncation of any of the IP attributes (i.e., ip, ipv6, xff), where 0 = no, 1 = yes (e.g., from 1.2.3.4 to 1.2.3.0).
	//   Refer to https://tools.ietf.org/html/rfc6235#section-4.1.1 for more information on IP truncation.
	IPTr int8 `json:"iptr,omitempty"`

	// Attribute:
	//   carrier
	// Type:
	//   string
	// Definition:
	//   Carrier or ISP (e.g., “VERIZON”) using exchange curated string names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`

	// Attribute:
	//   mccmnc
	// Type:
	//   string
	// Definition:
	//   Mobile carrier as the concatenated MCC-MNC code (e.g., “310-005” identifies Verizon Wireless CDMA in the USA).
	//   Refer to https://en.wikipedia.org/wiki/Mobile_country_code for further information and references.
	//   Note that the dash between the MCC and MNC parts is required to remove parsing ambiguity.
	MCCMNC string `json:"mccmnc,omitempty"`

	// Attribute:
	//   mccmncsim
	// Type:
	//   string
	// Definition:
	//   MCC and MNC of the SIM card using the same format as mccmnc.
	//   When both values are available, a difference between them reveals that a user is roaming.
	MCCMNCSIM string `json:"mccmncsim,omitempty"`

	// Attribute:
	//   contype
	// Type:
	//   integer
	// Definition:
	//   Network connection type.
	//   Refer to List: Connection Types.
	ConType ConnectionType `json:"contype,omitempty"`

	// Attribute:
	//   geofetch
	// Type:
	//   integer
	// Definition:
	//   Indicates if the geolocation API will be available to JavaScript code running in display ad, where 0 = no, 1 = yes.
	GeoFetch int8 `json:"geofetch,omitempty"`

	// Attribute:
	//   geo
	// Type:
	//   object
	// Definition:
	//   Location of the device (i.e., typically the user's current location).
	//   Refer to Object: Geo.
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
