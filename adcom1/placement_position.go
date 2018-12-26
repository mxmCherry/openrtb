package adcom

// PlacementPosition represents placement positions as a relative measure of visibility or prominence.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type PlacementPosition int

// Placement positions.
const (
	PositionAboveFold  PlacementPosition = 1 // Above The Fold
	PositionLocked     PlacementPosition = 2 // Locked (i.e., fixed position)
	PositionBelowFold  PlacementPosition = 3 // Below The Fold
	PositionHeader     PlacementPosition = 4 // Header
	PositionFooter     PlacementPosition = 5 // Footer
	PositionSideBar    PlacementPosition = 6 // Sidebar
	PositionFullScreen PlacementPosition = 7 // Fullscreen
)
