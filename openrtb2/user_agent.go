package openrtb2

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/v16/adcom1"
)

// 3.2.29 Object: UserAgent
//
// This object contains a single user identifier provided as part of extended identifiers.
// The exchange should ensure that business agreements allow for the sending of this data.
type UserAgent struct {

	// Attribute:
	//   browsers
	// Type:
	//   object array; recommended
	// Description:
	//   Each BrandVersion object (see Section 3.2.30) identifies a browser or similar
	//   software component. Implementers should send brands and versions
	//   derived from the Sec-CH-UA-Full-Version-List header.
	Browsers []BrandVersion `json:"browsers,omitempty"`

	// Attribute:
	//   platform
	// Type:
	//   object; recommended
	// Description:
	//   A BrandVersion object (see Section 3.2.30) that identifies the user agent’s
	//   execution platform / OS. Implementers should send a brand derived from the
	//   Sec-CH-UA-Platform header, and version derived from the Sec-CH-UA-
	//   Platform-Version header.
	Platform *BrandVersion `json:"platform,omitempty"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Description:
	//   1 if the agent prefers a “mobile” version of the content, if available, i.e.
	//   optimized for small screens or touch input. 0 if the agent prefers the “desktop”
	//   or “full” content. Implementers should derive this value from the Sec-CH-UA-
	//   Mobile header.
	Mobile *int8 `json:"mobile,omitempty"`

	// Attribute:
	//   architecture
	// Type:
	//   string
	// Description:
	//   Device’s major binary architecture, e.g. “x86” or “arm”. Implementers should
	//   retrieve this value from the Sec-CH-UA-Arch header.
	Architecture string `json:"architecture,omitempty"`

	// Attribute:
	//   bitness
	// Type:
	//   string
	// Description:
	//   Device’s bitness, e.g. “64” for 64-bit architecture. Implementers should
	//   retrieve this value from the Sec-CH-UA-Bitness header.
	Bitness string `json:"bitness,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model. Implementers should retrieve this value from the Sec-CH-UA-
	//   Model header.
	Model string `json:"model,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   integer; default 0
	// Description:
	//   The source of data used to create this object, List: User-Agent Source in
	//   AdCOM 1.0.
	Source adcom1.UserAgentSource `json:"source,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
