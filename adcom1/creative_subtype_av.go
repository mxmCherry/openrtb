package adcom1

// CreativeSubtypeAV represents subtypes of audio and video ad creatives.
type CreativeSubtypeAV int8 // TODO: rename to smth like AudioVideoCreativeSubtype? (sounds more natural)

// Subtypes of audio and video ad creatives.
const (
	CreativeVAST10         CreativeSubtypeAV = 1  // VAST 1.0
	CreativeVAST20         CreativeSubtypeAV = 2  // VAST 2.0
	CreativeVAST30         CreativeSubtypeAV = 3  // VAST 3.0
	CreativeVAST10Wrapper  CreativeSubtypeAV = 4  // VAST 1.0 Wrapper
	CreativeVAST20Wrapper  CreativeSubtypeAV = 5  // VAST 2.0 Wrapper
	CreativeVAST30Wrapper  CreativeSubtypeAV = 6  // VAST 3.0 Wrapper
	CreativeVAST40         CreativeSubtypeAV = 7  // VAST 4.0
	CreativeVAST40Wrapper  CreativeSubtypeAV = 8  // VAST 4.0 Wrapper
	CreativeDAAST10        CreativeSubtypeAV = 9  // DAAST 1.0
	CreativeDAAST10Wrapper CreativeSubtypeAV = 10 // DAAST 1.0 Wrapper
	CreativeVAST41         CreativeSubtypeAV = 11 // VAST 4.1
	CreativeVAST41Wrapper  CreativeSubtypeAV = 12 // VAST 4.1 Wrapper
)
