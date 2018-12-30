package adcom1

// CompanionType represents options to indicate markup types allowed for companion ads that apply to video and audio ads.
// This table is derived from VAST 2.0+ and DAAST 1.0+ specifications.
type CompanionType int8

// options to indicate markup types allowed for companion ads that apply to video and audio ads.
// This table is derived from VAST 2.0+ and DAAST 1.0+ specifications.
const (
	CompanionStatic CompanionType = 1 // Static Resource
	CompanionHTML   CompanionType = 2 // HTML Resource
	CompanionIFrame CompanionType = 3 // iframe Resource
)
