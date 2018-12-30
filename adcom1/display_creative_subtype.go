package adcom1

// DisplayCreativeSubtype represents subtypes of display ad creatives.
type DisplayCreativeSubtype int8

// Subtypes of display ad creatives.
const (
	CreativeHTML   DisplayCreativeSubtype = 1 // HTML
	CreativeAMP    DisplayCreativeSubtype = 2 // AMPHTML
	CreativeImage  DisplayCreativeSubtype = 3 // Structured Image Object
	CreativeNative DisplayCreativeSubtype = 4 // Structured Native Object
)
