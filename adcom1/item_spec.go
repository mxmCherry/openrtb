package adcom1

// ItemSpec represents a bunch of data for OpenRTB 3 Item.spec field.
// Refer to OpenRTB 3.0 specification for details.
//
//   For AdCOM v1.x, the objects allowed here are Placement and any objects subordinate to these as specified by AdCOM.
//
type ItemSpec struct {
	Placement *Placement `json:"placement,omitempty"`
}
