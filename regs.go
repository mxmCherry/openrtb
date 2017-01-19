package openrtb

// 3.2.16 Object: Regs
//
// This object contains any legal, governmental, or industry regulations that apply to the request. The coppa flag
// signals whether or not the request falls under the  United States Federal Trade Commission’s regulations for the
// United States Children’s Online Privacy Protection Act (“COPPA”).  Refer to Section 7.1 for more information.
type Regs struct {

	// Attribute:
	//   coppa
	// Type:
	//   integer
	// Description:
	//   Flag indicating if this request is subject to the COPPA
	//   regulations established by the USA FTC, where 0 = no, 1 = yes.
	COPPA int8 `json:"coppa,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext RawJSON `json:"ext,omitempty"`
}
