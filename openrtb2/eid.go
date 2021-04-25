package openrtb2

import "encoding/json"

// Eid defines the contract for bidrequest.user.eids
// Responsible for the Universal User ID support: establishing pseudonymous IDs for users.
type Eid struct {
	Source string          `json:"source"`
	ID     string          `json:"id,omitempty"`
	Uids   []Uid           `json:"uids,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}
