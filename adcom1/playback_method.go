package adcom1

// PlaybackMethod represents media playback methods.
type PlaybackMethod int8

// Media playback methods.
const (
	PlaybackPageLoadSoundOn  PlaybackMethod = 1 // Initiates on Page Load with Sound On
	PlaybackPageLoadSoundOff PlaybackMethod = 2 // Initiates on Page Load with Sound Off by Default
	PlaybackClickSoundOn     PlaybackMethod = 3 // Initiates on Click with Sound On
	PlaybackMouseOverSoundOn PlaybackMethod = 4 // Initiates on Mouse-Over with Sound On
	PlaybackViewportSoundOn  PlaybackMethod = 5 // Initiates on Entering Viewport with Sound On
	PlaybackViewportSoundOff PlaybackMethod = 6 // Initiates on Entering Viewport with Sound Off by Default
)
