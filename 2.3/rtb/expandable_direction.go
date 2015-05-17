package rtb

// 5.5 Expandable Direction
//
// The following table lists the directions in which an expandable ad may expand, given the positioning of
// the ad unit on the page and constraints imposed by the content.
const (
	ExpandableDirectionLeft       uint8 = 1 // 1 Left
	ExpandableDirectionRight      uint8 = 2 // 2 Right
	ExpandableDirectionUp         uint8 = 3 // 3 Up
	ExpandableDirectionDown       uint8 = 4 // 4 Down
	ExpandableDirectionFullScreen uint8 = 5 // 5 Full Screen
)
