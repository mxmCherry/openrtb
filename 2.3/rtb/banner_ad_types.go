package rtb

// 5.2 Banner Ad Types
//
// The following table indicates the types of ads that can be accepted by the exchange unless restricted by
// publisher site settings.
const (
	BannerAdTypeXHTMLText   uint8 = 1 // 1 XHTML Text Ad (usually mobile)
	BannerAdTypeXHTMLBanner uint8 = 2 // 2 XHTML Banner Ad. (usually mobile)
	BannerAdTypeJavaScript  uint8 = 3 // 3 JavaScript Ad; must be valid XHTML (i.e., Script Tags Included)
	BannerAdTypeIframe      uint8 = 4 // 4 iframe
)
