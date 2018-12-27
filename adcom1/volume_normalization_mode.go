package adcom

// VolumeNormalizationMode represents types of volume normalization modes, typically for audio.
type VolumeNormalizationMode int

// Types of volume normalization modes, typically for audio.
const (
	VolumeNormNone     VolumeNormalizationMode = 0 // None
	VolumeNormAvg      VolumeNormalizationMode = 1 // Ad Volume Average Normalized to Content
	VolumeNormPeak     VolumeNormalizationMode = 2 // Ad Volume Peak Normalized to Content
	VolumeNormLoudNorm VolumeNormalizationMode = 3 // Ad Loudness Normalized to Content
	VolumeNormCustom   VolumeNormalizationMode = 4 // Custom Volume Normalization
)
