package adcom

import "encoding/json"

// DisplayPlacement  object signals that the placement may be a display placement.
// It provides additional detail about permitted display ads including simple banners, AMPHTML (i.e., Accelerated Mobile Pages), and native.
type DisplayPlacement struct {
	// Attribute:
	//   pos
	// Type:
	//   integer
	// Definition:
	//   Placement position on screen.
	//   Refer to List: Placement Positions.
	Pos interface{} `json:"pos,omitempty"`

	// Attribute:
	//   instl
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicates if this is an interstitial placement, where 0 = no, 1 = yes.
	Instl int8 `json:"instl,omitempty"`

	// Attribute:
	//   topframe
	// Type:
	//   integer
	// Definition:
	//   Indicates if the placement will be loaded into an iframe or not, where 0 = unfriendly iframe or unknown, 1 = top frame, friendly iframe, or SafeFrame.
	//   A value of "1" can be understood to mean that expandable ads are technically capable of being delivered.
	TopFrame int8 `json:"topframe,omitempty"`

	// Attribute:
	//   ifrbust
	// Type:
	//   string array
	// Definition:
	//   Array of iframe busters supported by this placement.
	//   The meaning of strings in this attribute must be coordinated a priori among vendors.
	IfrBust []string `json:"ifrbust,omitempty"`

	// Attribute:
	//   clktype
	// Type:
	//   integer; default 1
	// Definition:
	//   Indicates the click type of this placement.
	//   Refer to List: Click Types.
	ClkType interface{} `json:"clktype,omitempty"`

	// Attribute:
	//   ampren
	// Type:
	//   integer
	// Definition:
	//   AMPHTML rendering treatment for AMP ads in this placement, where 1 = early loading, 2 = standard loading.
	AMPRen int8 `json:"ampren,omitempty"`

	// Attribute:
	//   ptype
	// Type:
	//   Integer; recommended
	// Definition:
	//   The display placement type.
	//   Refer to List: Display Placement Types.
	PType interface{} `json:"ptype,omitempty"`

	// Attribute:
	//   context
	// Type:
	//   integer; recommended
	// Definition:
	//   The context of the placement.
	//   Refer to List: Display Context Types.
	Context interface{} `json:"context,omitempty"`

	// Attribute:
	//   mime
	// Type:
	//   string array
	// Definition:
	//   Array of supported mime types (e.g., “image/jpeg”, “image/gif”).
	//   If omitted, all types are assumed.
	MIME []string `json:"mime,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   List of supported APIs.
	//   If an API is not explicitly listed, it is assumed to be unsupported.
	//   Refer to List: API Frameworks.
	API []interface{} `json:"api,omitempty"`

	// Attribute:
	//   ctype
	// Type:
	//   integer array
	// Definition:
	//   Creative subtypes permitted.
	//   Refer to List: Creative Subtypes - Display.
	CType []interface{} `json:"ctype,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Definition:
	//   Width of the placement in units specified by unit.
	//   Note that this size applies to the placement itself; permitted creative sizes are specified elsewhere (e.g., DisplayFormat, ImageAssetFormat, etc.).
	W int `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Definition:
	//   Width of the placement in units specified by unit.
	//   Note that this size applies to the placement itself; permitted creative sizes are specified elsewhere (e.g., DisplayFormat, ImageAssetFormat, etc.).
	H int `json:"h,omitempty"`

	// Attribute:
	//   unit
	// Type:
	//   integer; default 1
	// Definition:
	//   Unit of size used for placement size (i.e., w and h attributes).
	//   Refer to List: Size Units.
	Unit interface{} `json:"unit,omitempty"`

	// Attribute:
	//   priv
	// Type:
	//   integer; default 0
	// Definition:
	//   Indicator of whether or not the placement supports a buyer-specific privacy notice URL, where 0 = no, 1 = yes.
	Priv int8 `json:"priv,omitempty"`

	// Attribute:
	//   displayfmt
	// Type:
	//   object array
	// Definition:
	//   Array of objects that govern the attributes (e.g., sizes) of a banner display placement.
	//   Refer to Object: DisplayFormat.
	DisplayFmt []interface{} `json:"displayfmt,omitempty"`

	// Attribute:
	//   nativefmt
	// Type:
	//   object
	// Definition:
	//   This object specified the required and permitted assets and attributes of a native display placement.
	//   Refer to Object: NativeFormat.
	NativeFmt interface{} `json:"nativefmt,omitempty"`

	// Attribute:
	//   event
	// Type:
	//   object array
	// Definition:
	//   Array of supported ad tracking events.
	//   Refer to Object: EventSpec.
	Event []interface{} `json:"event,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
