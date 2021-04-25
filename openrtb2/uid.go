package openrtb2

import "encoding/json"

// Uid defines the contract for bidrequest.user.eids[i].uids[j]
type Uid struct {
	ID    string          `json:"id"`
	AType int             `json:"atype,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}
