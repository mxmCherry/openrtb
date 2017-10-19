package openrtb

// PtrInt8 returns pointer to passed argument.
func PtrInt8(n int8) *int8 {
	return &n
}

// PtrUint64 returns pointer to passed argument.
func PtrUint64(n uint64) *uint64 {
	return &n
}
