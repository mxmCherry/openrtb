package adcom1

// AgentType identifies the user agent types a user identifier is from.
type AgentType int

// Agent types dewscribing where the user agent is from.
//
// Values of 500+ hold vendor-specific codes.
const (
	ATypeWeb    AgentType = 1 // An ID which is tied to a specific web browser or device (cookie-based, probabilistic, or other).
	ATypeApp    AgentType = 2 // In-app impressions, which will typically contain a type of device ID (or rather, the privacy-compliant versions of device IDs).
	ATypePerson AgentType = 3 // A person-based ID, i.e., that is the same across devices.
)
