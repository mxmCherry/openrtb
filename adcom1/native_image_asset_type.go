package adcom

// NativeImageAssetType represents image asset types.
// This list is non-exhaustive and is intended to be expanded over time.
// Size recommendations are noted as "maximum height or width of at least", which means the publisher or supply platform should support a maximum height or width of at least this value and the buying platform knows that an image of this size should be accepted.
type NativeImageAssetType int

// Common image asset types.
//
// Values 500+ hold vendor-specific codes.
const (
	// Icon: Icon image. Maximum height at least 50 device independent pixels (DIPS); aspect ratio 1:1.
	ImageAssetIcon NativeImageAssetType = 1

	// Main: Large image preview for the ad. At least one of 2 size variants required:
	//
	// Small: Maximum height at least 627 DIPS; maximum width at least 627, 836, or 1198 DIPS (i.e., aspect ratios of 1:1, 4:3, or 1.91:1, respectively).
	//
	// Large: Maximum height at least 200 DIPS; maximum width at least 200, 267, or 382 DIPS (i.e., aspect ratios of 1:1, 4:3, or 1.91:1, respectively).
	ImageAssetMain NativeImageAssetType = 3
)
