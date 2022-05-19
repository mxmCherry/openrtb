package openrtb2

// MarkupType defines enum for creative markup type (Bid.mtype property).
type MarkupType int8

const (
	MarkupBanner = 1
	MarkupVideo  = 2
	MarkupAudio  = 3
	MarkupNative = 4
)
