package rtb

// 3.2.3 Object: Banner
//
// This object represents the most general type of impression. Although the term “banner” may have very
// specific meaning in other contexts, here it can be many things including a simple static image, an
// expandable ad unit, or even in-banner video (refer to the Video object in Section 3.2.4 for the more
// generalized and full featured video ad units). An array of Banner objects can also appear within the
// Video to describe optional companion ads defined in the VAST specification.
//
// The presence of a Banner as a subordinate of the Imp object indicates that this impression is offered as
// a banner type impression. At the publisher’s discretion, that same impression may also be offered as
// video and/or native by also including as Imp subordinates the Video and/or Native objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Banner struct {

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Width of the impression in pixels.
	//   If neither wmin nor wmax are specified, this value is an exact
	//   width requirement. Otherwise it is a preferred width.
	W uint64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Height of the impression in pixels.
	//   If neither hmin nor hmax are specified, this value is an exact
	//   height requirement. Otherwise it is a preferred height.
	H uint64 `json:"h,omitempty"`

	// Attribute:
	//   wmax
	// Type:
	//   integer
	// Description:
	//   Maximum width of the impression in pixels.
	//   If included along with a w value then w should be interpreted
	//   as a recommended or preferred width.
	WMax uint64 `json:"wmax,omitempty"`

	// Attribute:
	//   hmax
	// Type:
	//   integer
	// Description:
	//   Maximum height of the impression in pixels.
	//   If included along with an h value then h should be interpreted
	//   as a recommended or preferred height.
	HMax uint64 `json:"hmax,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer
	// Description:
	//   Minimum width of the impression in pixels.
	//   If included along with a w value then w should be interpreted
	//   as a recommended or preferred width.
	WMin uint64 `json:"wmin,omitempty"`

	// Attribute:
	//   hmin
	// Type:
	//   integer
	// Description:
	//   Minimum height of the impression in pixels.
	//   If included along with an h value then h should be interpreted
	//   as a recommended or preferred height.
	HMin uint64 `json:"hmin,omitempty"`

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Unique identifier for this banner object. Recommended when
	//   Banner objects are used with a Video object (Section 3.2.4) to
	//   represent an array of companion ads. Values usually start at 1
	//   and increase with each object; should be unique within an
	//   impression.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   btype
	// Type:
	//   integer array
	// Description:
	//   Blocked banner ad types. Refer to List 5.2.
	BType []uint8 `json:"btype,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BAttr []uint8 `json:"battr,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List 5.4
	Pos uint8 `json:"pos,omitempty"`

	// Attribute:
	//   mimes
	// Type:
	//   string array
	// Description:
	//   Content MIME types supported. Popular MIME types may
	//   include “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	MIMEs []string `json:"mimes,omitempty"`

	// Attribute:
	//   topframe
	// Type:
	//   integer
	// Description:
	//    Indicates if the banner is in the top frame as opposed to an
	//    iframe, where 0 = no, 1 = yes.
	TopFrame uint8 `json:"topframe,omitempty"`

	// Attribute:
	//   expdir
	// Type:
	//   integer array
	// Description:
	//   Directions in which the banner may expand. Refer to List 5.5.
	ExpDir []uint8 `json:"expdir,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List 5.6. If an API is not explicitly listed, it is assumed not to be supported.
	API []uint8 `json:"api,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext,omitempty"`
}
