package adcom1

// NativeDataAssetType represents data asset types.
// This list is non-exhaustive and is intended to be expanded over time.
// Size recommendations are noted as "maximum length of at least", which means the publisher or supply platform should support a maximum length of at least this value and the buying platform knows that a string of this size should be accepted.
type NativeDataAssetType int

// Common data asset types.
//
// Values of 500+ hold vendor-specific codes.
const (
	DataAssetSponsored  NativeDataAssetType = 1  // sponsored: "Sponsored By" message which should contain the brand name of the sponsor. Recommended maximum length of at least 25 characters.
	DataAssetDesc       NativeDataAssetType = 2  // desc: Descriptive text associated with the product or service being advertised. Long text lengths may be truncated or ellipsed when rendered. Recommended maximum length of at least 140 characters.
	DataAssetRating     NativeDataAssetType = 3  // rating: Numeric rating of the product (e.g., an app's rating). Recommended integer range of 0-5.
	DataAssetLikes      NativeDataAssetType = 4  // likes: Number of social ratings or "likes" of the product.
	DataAssetDownloads  NativeDataAssetType = 5  // downloads: Number downloads and/or installs of the product.
	DataAssetPrice      NativeDataAssetType = 6  // price: Price of the product, app, or in-app purchase. Value should include currency symbol in localized format (e.g., "$10").
	DataAssetSalePrice  NativeDataAssetType = 7  // saleprice: Sale price that can be used together with "price" to indicate a comparative discounted price. Value should include currency symbol in localized format (e.g., "$8.50").
	DataAssetPhone      NativeDataAssetType = 8  // phone: A formatted phone number.
	DataAssetAddress    NativeDataAssetType = 9  // address: A formatted address.
	DataAssetDesc2      NativeDataAssetType = 10 // desc2: Additional descriptive text associated with the product.
	DataAssetDisplayURL NativeDataAssetType = 11 // displayurl: Display URL for the ad. To be used when sponsoring entity doesn't own the content (e.g., "Sponsored by Brand on Site", where Site is specified in this data asset).
	DataAssetCTAText    NativeDataAssetType = 12 // ctatext: Description of the call to action (CTA) button for the destination URL. Recommended maximum length of at least 15 characters.
)
