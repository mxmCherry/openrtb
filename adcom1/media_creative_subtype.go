package adcom1

// MediaCreativeSubtype represents subtypes of audio and video ad creatives.
type MediaCreativeSubtype int8

// Subtypes of audio and video ad creatives.
const (
	CreativeVAST10         MediaCreativeSubtype = 1  // VAST 1.0
	CreativeVAST20         MediaCreativeSubtype = 2  // VAST 2.0
	CreativeVAST30         MediaCreativeSubtype = 3  // VAST 3.0
	CreativeVAST10Wrapper  MediaCreativeSubtype = 4  // VAST 1.0 Wrapper
	CreativeVAST20Wrapper  MediaCreativeSubtype = 5  // VAST 2.0 Wrapper
	CreativeVAST30Wrapper  MediaCreativeSubtype = 6  // VAST 3.0 Wrapper
	CreativeVAST40         MediaCreativeSubtype = 7  // VAST 4.0
	CreativeVAST40Wrapper  MediaCreativeSubtype = 8  // VAST 4.0 Wrapper
	CreativeDAAST10        MediaCreativeSubtype = 9  // DAAST 1.0
	CreativeDAAST10Wrapper MediaCreativeSubtype = 10 // DAAST 1.0 Wrapper
	CreativeVAST41         MediaCreativeSubtype = 11 // VAST 4.1
	CreativeVAST41Wrapper  MediaCreativeSubtype = 12 // VAST 4.1 Wrapper
	CreativeVAST42         MediaCreativeSubtype = 13 // VAST 4.2
	CreativeVAST42Wrapper  MediaCreativeSubtype = 14 // VAST 4.2 Wrapper
)
