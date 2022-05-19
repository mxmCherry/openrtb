package adcom1

// StartDelay represents video or audio start delay.
type StartDelay int64

// Options for the video or audio start delay.
// If the start delay value is greater than 0, then the position is mid-roll and the value indicates the start delay.
const (
	StartPreRoll  StartDelay = 0  // Pre-Roll
	StartMidRoll  StartDelay = -1 // Generic Mid-Roll
	StartPostRoll StartDelay = -2 // Generic Post-Roll
)

// Ptr returns pointer to own value.
func (d StartDelay) Ptr() *StartDelay {
	return &d
}

// Val safely dereferences pointer, returning default value (StartDelayPreRoll) for nil.
func (d *StartDelay) Val() StartDelay {
	if d == nil {
		return StartPreRoll
	}
	return *d
}
