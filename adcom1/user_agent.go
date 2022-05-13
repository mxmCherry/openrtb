package adcom1

import "encoding/json"

// UserAgent represents Structured user agent information provided when client supports User-Agent Client Hints.
// If both device.ua and device.sua are present in the bid request, device.sua should be considered the more accurate representation of the device attributes.
// This is because the device.ua may contain a frozen or reduced UserAgent string.
type UserAgent struct {
	// Attribute:
	//   browsers
	// Type:
	//   object array; recommended
	// Definition:
	//   Each BrandVersion object identifies a browser or similar software component.
	//   Refer to Object: BrandVersion.
	//   Implementers should send brands and versions derived from the Sec-CH-UA-Full-Version-List header.
	Browsers []BrandVersion `json:"browsers,omitempty"`

	// Attribute:
	//   platform
	// Type:
	//   object; recommended
	// Definition:
	//   Refer to Object: BrandVersion that identifies the user agent’s execution platform / OS.
	//   Implementers should send a brand derived from the Sec-CH-UA-Platform header, and version derived from the Sec-CH-UA-Platform-Version header.
	Platform BrandVersion `json:"platform,omitempty"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Definition:
	//   1 if the agent prefers a “mobile” version of the content, if available, i.e. optimized for small screens or touch input.
	//   0 if the agent prefers the “desktop” or “full” content.
	//   Implementers should derive this value from the Sec-CH-UA-Mobile header.
	Mobile int8 `json:"mobile,omitempty"`

	// Attribute:
	//   architecture
	// Type:
	//   string
	// Definition:
	//   Device’s major binary architecture, e.g. “x86” or “arm”.
	//   Implementers should retrieve this value from the Sec-CH-UA-Arch header.
	Architecture string `json:"architecture,omitempty"`

	// Attribute:
	//   bitness
	// Type:
	//   string
	// Definition:
	//   Device’s bitness, e.g. “64” for 64-bit architecture.
	//   Implementers should retrieve this value from the Sec-CH-UA-Bitness header.
	Bitness string `json:"bitness,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Definition:
	//   Device model.
	//   Implementers should retrieve this value from the Sec-CH-UA-Model header.
	Model string `json:"model,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   integer
	// Definition:
	//   The source of data used to create this object.
	//   Refer to List: User-Agent Source
	Source UserAgentSource `json:"source,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
