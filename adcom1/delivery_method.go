package adcom

// DeliveryMethod represents options for the delivery of video or audio content.
type DeliveryMethod int8

// Options for the delivery of video or audio content.
const (
	DeliveryStreaming   DeliveryMethod = 1 // Streaming
	DeliveryProgressive DeliveryMethod = 2 // Progressive
	DeliveryDownload    DeliveryMethod = 3 // Download
)
