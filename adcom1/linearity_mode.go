package adcom

// LinearityMode represents options for media linearity, typically for video.
type LinearityMode int8

// Options for media linearity, typically for video.
const (
	LinearityLinear    LinearityMode = 1 // Linear (i.e., In-Stream such as Pre-Roll, Mid-Roll, Post-Roll)
	LinearityNonLinear LinearityMode = 2 // Non-Linear (i.e., Overlay)
)
