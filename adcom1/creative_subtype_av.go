package adcom1

// CreativeSubtypeAV represents subtypes of audio and video ad creatives.
type CreativeSubtypeAV int8 // TODO: rename to smth like AudioVideoCreativeSubtype? (sounds more natural)

// Subtypes of audio and video ad creatives.
const (
	CreativeVAST1         CreativeSubtypeAV = 1  // VAST 1.0
	CreativeVAST2         CreativeSubtypeAV = 2  // VAST 2.0
	CreativeVAST3         CreativeSubtypeAV = 3  // VAST 3.0
	CreativeVAST1Wrapper  CreativeSubtypeAV = 4  // VAST 1.0 Wrapper
	CreativeVAST2Wrapper  CreativeSubtypeAV = 5  // VAST 2.0 Wrapper
	CreativeVAST3Wrapper  CreativeSubtypeAV = 6  // VAST 3.0 Wrapper
	CreativeVAST4         CreativeSubtypeAV = 7  // VAST 4.0
	CreativeVAST4Wrapper  CreativeSubtypeAV = 8  // VAST 4.0 Wrapper
	CreativeDAAST1        CreativeSubtypeAV = 9  // DAAST 1.0
	CreativeDAAST1Wrapper CreativeSubtypeAV = 10 // DAAST 1.0 Wrapper
	CreativeVAST41        CreativeSubtypeAV = 11 // VAST 4.1
	CreativeVAST41Wrapper CreativeSubtypeAV = 12 // VAST 4.1 Wrapper
)
