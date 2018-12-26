package adcom

// DeviceType represents types of devices.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type DeviceType int

// Types of devices.
const (
	DeviceMobile    DeviceType = 1 // Mobile/Tablet - General
	DevicePC        DeviceType = 2 // Personal Computer
	DeviceTV        DeviceType = 3 // Connected TV
	DevicePhone     DeviceType = 4 // Phone
	DeviceTablet    DeviceType = 5 // Tablet
	DeviceConnected DeviceType = 6 // Connected Device
	DeviceSetTopBox DeviceType = 7 // Set Top Box
)
