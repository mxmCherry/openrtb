package adcom1

// PlacementPosition represents placement positions as a relative measure of visibility or prominence.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type PlacementPosition int8

// Placement positions.
const (
	PositionUnknown    PlacementPosition = 0 // Unknown
	PositionAboveFold  PlacementPosition = 1 // Above The Fold
	PositionLocked     PlacementPosition = 2 // Locked (i.e., fixed position)
	PositionBelowFold  PlacementPosition = 3 // Below The Fold
	PositionHeader     PlacementPosition = 4 // Header
	PositionFooter     PlacementPosition = 5 // Footer
	PositionSideBar    PlacementPosition = 6 // Sidebar
	PositionFullScreen PlacementPosition = 7 // Fullscreen
)

// Ptr returns pointer to own value.
func (p PlacementPosition) Ptr() *PlacementPosition {
	return &p
}

// Val safely dereferences pointer, returning default value (AdPositionUnknown) for nil.
func (p *PlacementPosition) Val() PlacementPosition {
	if p == nil {
		return PositionUnknown
	}
	return *p
}
