package rtb

// Mobile-optimized signal, where 0 = no, 1 = yes.
const (
	SiteMobileNo  uint8 = 0 // 0 = no
	SiteMobileYes uint8 = 1 // 1 = yes
)

// Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
const (
	SitePrivacyPolicyNo  uint8 = 0 // 0 = no
	SitePrivacyPolicyYes uint8 = 1 // 1 = yes
)

// 3.2.6 Object: Site
// 
// This object should be included if the ad supported content is a website as opposed to a non-browser
// application. A bid request must not contain both a Site and an App object. At a minimum, it is useful
// to provide a site ID or page URL, but this is not strictly required.
type Site struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific site ID.
	ID string `json:"id"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Site name (may be aliased at the publisher’s request).
	Name string `json:"name"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the site (e.g., “mysite.foo.com”).
	Domain string `json:"domain"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories of the site. Refer to List 5.1.
	Cat []string `json:"cat"`

	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current
	//   section of the site. Refer to List 5.1.
	SectionCat []string `json:"sectioncat"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current page
	//   or view of the site. Refer to List 5.1.
	PageCat []string `json:"pagecat"`

	// Attribute:
	//   page
	// Type:
	//   string
	// Description:
	//   URL of the page where the impression will be shown.
	Page string `json:"page"`

	// Attribute:
	//   ref
	// Type:
	//   string
	// Description:
	//   Referrer URL that caused navigation to the current page.
	Ref string `json:"ref"`

	// Attribute:
	//   search
	// Type:
	//   string
	// Description:
	//   Search string that caused navigation to the current page.
	Search string `json:"search"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Description:
	//   Mobile-optimized signal, where 0 = no, 1 = yes.
	Mobile uint8 `json:"mobile"`

	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy uint8 `json:"privacypolicy"`

	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.8) of the site.
	Publisher Publisher `json:"publisher"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.9) within the site.
	Content Content `json:"content"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the site.
	Keywords string `json:"keywords"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}