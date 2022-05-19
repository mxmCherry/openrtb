package adcom1

// ProductionQuality represents content quality.
type ProductionQuality int8

// Options for content quality.
// These values are defined by the IAB; refer to www.iab.com/wp-content/uploads/2015/03/long-form-video-final.pdf for more information.
const (
	ProductionUnknown      ProductionQuality = 0 // Unknown
	ProductionProfessional ProductionQuality = 1 // Professionally Produced
	ProductionProsumer     ProductionQuality = 2 // Prosumer
	ProductionUser         ProductionQuality = 3 // User Generated (UGC)
)
