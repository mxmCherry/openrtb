package adcom

// SizeUnit represents units of height and width used by creatives, assets, and placement specifications where noted.
type SizeUnit int8

// Units of height and width used by creatives, assets, and placement specifications where noted.
const (
	SizeDIP SizeUnit = 1 // Device Independent Pixels (DIPS)
	SizeIn  SizeUnit = 2 // Inches
	SizeCm  SizeUnit = 3 // Centimeters
)
