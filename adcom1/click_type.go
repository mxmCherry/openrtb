package adcom

// ClickType represents types of creative activation (i.e., click) behavior types.
type ClickType int

// Types of creative activation (i.e., click) behavior types.
//
// Values of 500+ hold vendor-specific codes.
const (
	ClickNonClickable ClickType = 0 // Non-Clickable
	ClickUnknown      ClickType = 1 // Clickable - Details Unknown
	ClickEmbedded     ClickType = 2 // Clickable - Embedded Browser/Webview
	ClickNative       ClickType = 3 // Clickable - Native Browser
)
