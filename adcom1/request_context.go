package adcom1

// RequestContext represents a bunch of data for OpenRTB 3 Request.context field.
// Refer to OpenRTB 3.0 specification for details.
//
//   For AdCOM v1.x, the objects allowed here all of which are optional are one of the DistributionChannel subtypes (i.e., Site, App, or Dooh), User, Device, Regs, Restrictions, and any objects subordinate to these as specified by AdCOM.
//
type RequestContext struct {
	Site *Site `json:"site,omitempty"`
	App  *App  `json:"app,omitempty"`
	DOOH *DOOH `json:"dooh,omitempty"`

	User         *User         `json:"user,omitempty"`
	Device       *Device       `json:"device,omitempty"`
	Regs         *Regs         `json:"regs,omitempty"`
	Restrictions *Restrictions `json:"restrictions,omitempty"`
}
