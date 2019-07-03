package adcom1

// ExpandableDirection represents directions in which an expandable ad may expand, given the positioning of the ad unit on the page and constraints imposed by the content.
type ExpandableDirection int8

// Directions in which an expandable ad may expand.
const (
	ExpandableLeft       ExpandableDirection = 1 // Left
	ExpandableRight      ExpandableDirection = 2 // Right
	ExpandableUp         ExpandableDirection = 3 // Up
	ExpandableDown       ExpandableDirection = 4 // Down
	ExpandableFullScreen ExpandableDirection = 5 // Full Screen
)
