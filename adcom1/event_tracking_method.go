package adcom

// EventTrackingMethod represents methods of tracking of ad events.
// Vendor specific codes may include custom measurement companies (e.g., Moat, Doubleverify, IAS, etc.).
type EventTrackingMethod int

// Available methods of tracking of ad events.
//
// Values of 500+ hold vendor-specific codes.
const (
	TrackingImagePixel EventTrackingMethod = 1 // Image-Pixel: URL provided will be inserted as a 1x1 pixel at the time of the event.
	TrackingJS         EventTrackingMethod = 2 // JavaScript: URL provided will be inserted as a JavaScript tag at the time of the event.
)
