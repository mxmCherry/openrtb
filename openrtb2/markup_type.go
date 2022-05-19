package openrtb2

// MarkupType defines the type of the creative markup so that it can properly be
// associated with the right sub-object of the BidRequest.Imp.
//
// Originates from Bid.mtype property, not a separately-defined enum.
type MarkupType int8

const (
	MarkupBanner MarkupType = 1
	MarkupVideo  MarkupType = 2
	MarkupAudio  MarkupType = 3
	MarkupNative MarkupType = 4
)
