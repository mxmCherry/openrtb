package openrtb

// AuctionType defines an auction type.
type AuctionType int64

// AuctionType values.
const (
	FirstPrice      AuctionType = 1
	SecondPricePlus             = 2
	DealPrice                   = 3 // the value passed in flr is the agreed upon deal price

	// Values greater than 500 can be used for exchange-specific auction types.
)
