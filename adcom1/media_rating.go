package adcom

// MediaRating represents media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization.
// Refer to www.iab.com/guidelines/digital-video-suite for more information.
type MediaRating int8

// Media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization.
const (
	MediaRatingAll    MediaRating = 1 // All Audiences
	MediaRatingOver12 MediaRating = 2 // Everyone Over Age 12
	MediaRatingMature MediaRating = 3 // Mature Audiences
)
