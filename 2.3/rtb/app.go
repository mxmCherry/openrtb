package rtb

// 3.2.7 Object: App
//
// This object should be included if the ad supported content is a non-browser application (typically in
// mobile) as opposed to a website. A bid request must not contain both an App and a Site object. At a
// minimum, it is useful to provide an App ID or bundle, but this is not strictly required.
type App struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific app ID.
	ID string `json:"id"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   App name (may be aliased at the publisher’s request).
	Name string `json:"name"`

	// Attribute:
	//   bundle
	// Type:
	//   string
	// Description:
	//   Application bundle or package name (e.g., com.foo.mygame);
	//   intended to be a unique ID across exchanges.
	Bundle string `json:"bundle"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain"`

	// Attribute:
	//   storeurl
	// Type:
	//   string
	// Description:
	//   App store URL for an installed app; for QAG 1.5 compliance.
	StoreURL string `json:"storeurl"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories of the app. Refer to List 5.1.
	Cat []string `json:"cat"`

	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current
	//   section of the app. Refer to List 5.1.
	SectionCat []string `json:"sectioncat"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current page
	//   or view of the app. Refer to List 5.1.
	PageCat []string `json:"pagecat"`

	// Attribute:
	//   ver
	// Type:
	//   string
	// Description:
	//   Application version.
	Ver string `json:"ver"`

	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy uint8 `json:"privacypolicy"`

	// Attribute:
	//   paid
	// Type:
	//   integer
	// Description:
	//   0 = app is free, 1 = the app is a paid version.
	Paid uint8 `json:"paid"`

	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.8) of the app.
	Publisher Publisher `json:"publisher"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.9) within the app.
	Content Content `json:"content"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the app.
	Keywords string `json:"keywords"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}
