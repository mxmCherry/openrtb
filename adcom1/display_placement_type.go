package adcom

// DisplayPlacementType represents types of display placements; the locations where a native ad may be shown in relationship to the surrounding content.
type DisplayPlacementType int

// General types of display placements.
//
// Values of 500+ hold vendor-specific codes.
const (
	DisplayPlacementFeed    DisplayPlacementType = 1 // In the feed of content (e.g., as an item inside the organic feed, grid, listing, carousel, etc.).
	DisplayPlacementUnit    DisplayPlacementType = 2 // In the atomic unit of the content (e.g., in the article page or single image page).
	DisplayPlacementOutside DisplayPlacementType = 3 // Outside the core content (e.g., in the ads section on the right rail, as a banner-style placement near the content, etc.).
	DisplayPlacementWidget  DisplayPlacementType = 4 // Recommendation widget; most commonly presented below article content.
)
