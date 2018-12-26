package adcom

// LocationType represents options to indicate how the geographic information was determined.
type LocationType int

// Options to indicate how the geographic information was determined.
const (
	LocationGPS          LocationType = 1 // GPS/Location Services
	LocationIP           LocationType = 2 // IP Address
	LocationUserProvided LocationType = 3 // User Provided (e.g., registration data)
)
