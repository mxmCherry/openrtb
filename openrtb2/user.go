package openrtb2

import "encoding/json"

// 3.2.20 Object: User
//
// This object contains information known or derived about the human user of the device (i.e., the audience for advertising).
// The user id is an exchange artifact and may be subject to rotation or other privacy policies.
// However, this user ID must be stable long enough to serve reasonably as the basis for frequency capping and retargeting.
type User struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific ID for the user. At least one of id or
	//   buyeruid is recommended.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   buyeruid
	// Type:
	//   string; recommended
	// Description:
	//   Buyer-specific ID for the user as mapped by the exchange for
	//   the buyer. At least one of buyeruid or id is recommended.
	BuyerUID string `json:"buyeruid,omitempty"`

	// Attribute:
	//   yob
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   Year of birth as a 4-digit integer.
	Yob int64 `json:"yob,omitempty"`

	// Attribute:
	//   gender
	// Type:
	//   string; DEPRECATED
	// Description:
	//   Gender, where “M” = male, “F” = female, “O” = known to be
	//   other (i.e., omitted is unknown).
	Gender string `json:"gender,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords, interests, or intent. Only
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
	//   customdata
	// Type:
	//   string
	// Description:
	//   Optional feature to pass bidder data that was set in the
	//   exchange’s cookie. The string must be in base85 cookie safe
	//   characters and be in any format. Proper JSON encoding must
	//   be used to include “escaped” quotation marks.
	CustomData string `json:"customdata,omitempty"`

	// Attribute:
	//   geo
	// Type
	//   object
	// Description:
	//   Location of the user’s home base defined by a Geo object
	//   (Section 3.2.19). This is not necessarily their current location.
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   data
	// Type:
	//   object array
	// Description:
	//   Additional user data. Each Data object (Section 3.2.21)
	//   represents a different data source.
	Data []Data `json:"data,omitempty"`

	// Attribute:
	//   consent
	// Type:
	//   string
	// Description:
	//   When GDPR regulations are in effect this attribute contains
	//   the Transparency and Consent Framework’s Consent String
	//   data structure.
	Consent string `json:"consent,omitempty"`

	// Attribute:
	//   eids
	// Type:
	//   object array
	// Description:
	//   Details for support of a standard protocol for multiple third
	//   party identity providers (Section 3.2.27)
	EIDs []EID `json:"eids,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
