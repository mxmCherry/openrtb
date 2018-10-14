package openrtb3

import (
	"bytes"
)

// Bool is a bool type wrapper.
type Bool bool

// Val is a nil-safe value getter.
func (b *Bool) Val(defaultValue bool) bool {
	if b == nil {
		return defaultValue
	}
	return (bool)(*b)
}

// Ptr is a value-copying pointer shortcut.
//
// Intended usage:
//
//   req := Request {
//     WSeat: Bool(true).Ptr(),
//   }
//
func (b Bool) Ptr() *Bool {
	return &b
}

// MarshalJSON marshals Bool as JSON.
func (b Bool) MarshalJSON() (bb []byte, err error) {
	if b {
		bb = []byte("1")
	} else {
		bb = []byte("0")
	}
	return bb, err
}

// UnmarshalJSON unmarshals Bool from JSON.
func (b *Bool) UnmarshalJSON(bb []byte) error {
	if bytes.Equal(bb, []byte("1")) {
		*b = true
	} else {
		*b = false
	}
	return nil
}
