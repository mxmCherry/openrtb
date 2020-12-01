package openrtb2

type Eid struct {
	Source string `json:"source"`
	Uids   []Uid  `json:"uids"`
}

type Uid struct {
	ID    string `json:"id"`
	Atype int    `json:"atype,omitempty"`
}
