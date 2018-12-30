package adcom1

// BidMedia represents a bunch of data for OpenRTB 3 Bid.media field.
// Refer to OpenRTB 3.0 specification for details.
//
//   For AdCOM v1.x, the objects allowed here are “Ad” and any objects subordinate thereto as specified by AdCOM.
//
type BidMedia struct {
	Ad *Ad `json:"ad,omitempty"`
}
