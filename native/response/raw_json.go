package response

import "errors"

// RawJSON is a raw encoded JSON value.
// It implements encoding/json.Marshaler and encoding/json.Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
//
// Basically, it's just a copy of encoding/json.RawMessage type,
// but with more convenient non-pointer encoding.
//
// HEADS UP: this will be replaced with json.RawMessage when Go 1.10 is out.
type RawJSON []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawJSON) MarshalJSON() ([]byte, error) {
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawJSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("openrtb/native/response.RawJSON: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
