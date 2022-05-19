package adcom1

// StartDelayMode represents video or audio start delay.
type StartDelayMode int64

// Options for the video or audio start delay.
// If the start delay value is greater than 0, then the position is mid-roll and the value indicates the start delay.
const (
	StartDelayPreRoll         StartDelayMode = 0  // Pre-Roll
	StartDelayGenericMidRoll  StartDelayMode = -1 // Generic Mid-Roll
	StartDelayGenericPostRoll StartDelayMode = -2 // Generic Post-Roll
)

// Ptr returns pointer to own value.
func (s StartDelayMode) Ptr() *StartDelayMode {
	return &s
}

// Val safely dereferences pointer, returning default value (StartDelayPreRoll) for nil.
func (s *StartDelayMode) Val() StartDelayMode {
	if s == nil {
		return StartDelayPreRoll
	}
	return *s
}
