package rtb

// 5.4 Ad Position
//
// The following table specifies the position of the ad as a relative measure of visibility or prominence.
//
// This OpenRTB table has values derived from the IAB Quality Assurance Guidelines (QAG). Practitioners
// should keep in sync with updates to the QAG values as published on IAB.net. Values “3” – “6” apply to
// apps per the mobile addendum to QAG version 1.5.
//
// TODO: review and rename
const (
	AdPositionUnknown               uint8 = 0 // 0 Unknown
	AdPositionAboveTheFold          uint8 = 1 // 1 Above the Fold
	AdPositionMayBeInitiallyVisible uint8 = 2 // 2 DEPRECATED - May or may not be initially visible depending on screen size/resolution.
	AdPositionBelowTheFold          uint8 = 3 // 3 Below the Fold
	AdPositionHeader                uint8 = 4 // 4 Header
	AdPositionFooter                uint8 = 5 // 5 Footer
	AdPositionSidebar               uint8 = 6 // 6 Sidebar
	AdPositionFullScreen            uint8 = 7 // 7 Full Screen
)
