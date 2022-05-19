package adcom1

// VolumeNormalizationMode represents types of volume normalization modes, typically for audio.
type VolumeNormalizationMode int8

// Types of volume normalization modes, typically for audio.
const (
	VolumeNormNone     VolumeNormalizationMode = 0 // None
	VolumeNormAvg      VolumeNormalizationMode = 1 // Ad Volume Average Normalized to Content
	VolumeNormPeak     VolumeNormalizationMode = 2 // Ad Volume Peak Normalized to Content
	VolumeNormLoudness VolumeNormalizationMode = 3 // Ad Loudness Normalized to Content
	VolumeNormCustom   VolumeNormalizationMode = 4 // Custom Volume Normalization
)

// Ptr returns pointer to own value.
func (v VolumeNormalizationMode) Ptr() *VolumeNormalizationMode {
	return &v
}

// Val safely dereferences pointer, returning default value (VolumeNormNone) for nil.
func (v *VolumeNormalizationMode) Val() VolumeNormalizationMode {
	if v == nil {
		return VolumeNormNone
	}
	return *v
}
