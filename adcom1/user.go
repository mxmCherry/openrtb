package adcom1

import "encoding/json"

// User object contains information known or derived about the human user of the device (i.e., the audience for advertising).
// The user ID is a vendor-specific artifact and may be subject to rotation or other privacy policies.
// However, this user ID must be stable long enough to serve reasonably as the basis for frequency capping and retargeting.
type User struct {
	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Definition:
	//   Vendor-specific ID for the user.
	//   At least one of id or buyeruid is strongly recommended.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   buyeruid
	// Type:
	//   string; recommended
	// Definition:
	//   Buyer-specific ID for the user as mapped by an exchange for the buyer.
	//   At least one of id or buyeruid is strongly recommended.
	BuyerUID string `json:"buyeruid,omitempty"`

	// Attribute:
	//   yob
	// Type:
	//   integer; DEPRECATED
	// Definition:
	//   Year of birth as a 4-digit integer.
	YOB int64 `json:"yob,omitempty"`

	// Attribute:
	//   gender
	// Type:
	//   string; DEPRECATED
	// Definition:
	//   Gender, where “M” = male, “F” = female, “O” = known to be other (i.e., omitted is unknown).
	Gender string `json:"gender,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string; DEPRECATED
	// Definition:
	//   Comma-separated list of keywords, interests, or intent.
	//   Only one of 'keywords' or 'kwarray' may be present.
	//   NOTE: this field is deprecated, use 'kwarray' instead.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   kwarray
	// Type:
	//   string array
	// Definition:
	//   Array of keywords about the site. Only one of 'keywords' or 'kwarray' may be present.
	KwArray []string `json:"kwarray,omitempty"`

	// Attribute:
	//   consent
	// Type:
	//   string
	// Definition:
	//   GDPR consent string if applicable, complying with the comply with the IAB standard Consent String Format in the Transparency and Consent Framework technical specifications.
	Consent string `json:"consent,omitempty"`

	// Attribute:
	//   geo
	// Type:
	//   object
	// Definition:
	//   Location of the user's home base (i.e., not necessarily their current location).
	//   Refer to Object: Geo.
	Geo *Geo `json:"geo,omitempty"`

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
	//   eids
	// Type:
	//   object array
	// Definition:
	//   Extended (third-party) identifiers for this user. Refer to Object: Extended Identifiers.
	EIDs []ExtendedIdentifier `json:"eids,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
