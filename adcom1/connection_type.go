package adcom1

// ConnectionType represents options for the type of device connectivity.
type ConnectionType int8

// Options for the type of device connectivity.
const (
	ConnectionEthernet ConnectionType = 1 // 1	Ethernet; Wired Connection
	ConnectionWIFI     ConnectionType = 2 // 2	WIFI
	ConnectionCellular ConnectionType = 3 // 3	Cellular Network - Unknown Generation
	Connection2G       ConnectionType = 4 // 4	Cellular Network - 2G
	Connection3G       ConnectionType = 5 // 5	Cellular Network - 3G
	Connection4G       ConnectionType = 6 // 6	Cellular Network - 4G
	Connection5G       ConnectionType = 7 // 7	Cellular Network - 5G
)
