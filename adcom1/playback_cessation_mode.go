package adcom1

// PlaybackCessationMode represents modes for when media playback terminates.
type PlaybackCessationMode int8

// Modes for when media playback terminates.
const (
	PlaybackCompletion      PlaybackCessationMode = 1 // On Video Completion or when Terminated by User
	PlaybackLeavingViewport PlaybackCessationMode = 2 // On Leaving Viewport or when Terminated by User
	PlaybackFloating        PlaybackCessationMode = 3 // On Leaving Viewport Continues as a Floating/Slider Unit until Video Completion or when Terminated by User
)
