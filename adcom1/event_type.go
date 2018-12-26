package adcom

// EventType represents types of ad events available for tracking.
// These types refer to the actual event, timing, etc.; not the method of firing.
// Scripts that are performing measurement should be deployed at the "loaded" event.
type EventType int

// Types of ad events available for tracking
//
// Values of 500+ hold vendor-specific codes.
const (
	EventLoaded      EventType = 1 // loaded: Delivered as a part of the creative markup. Creative may be pre-cached or pre-loaded; prior to initial rendering.
	EventImpression  EventType = 2 // impression: Ad impression per IAB/MRC Ad Impression Measurement Guidelines.
	EventViewMRC50   EventType = 3 // viewable-mrc50: Visible impression using MRC definition of 50% in view for 1 second.
	EventViewMRC100  EventType = 4 // viewable-mrc100: 100% in view for 1 second (i.e., the GroupM standard).
	EventViewVideo50 EventType = 5 // viewable-video50: Visible impression for video using MRC definition of 50% in view for 2 seconds.
)
