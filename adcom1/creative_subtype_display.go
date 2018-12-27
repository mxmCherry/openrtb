package adcom

// CreativeSubtypeDisplay represents subtypes of display ad creatives.
type CreativeSubtypeDisplay int8

// Subtypes of display ad creatives.
const (
	CreativeHTML   CreativeSubtypeDisplay = 1 // HTML
	CreativeAMP    CreativeSubtypeDisplay = 2 // AMPHTML
	CreativeImage  CreativeSubtypeDisplay = 3 // Structured Image Object
	CreativeNative CreativeSubtypeDisplay = 4 // Structured Native Object
)
