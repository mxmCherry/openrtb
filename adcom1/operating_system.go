package adcom1

// OperatingSystem represents device operating system.
type OperatingSystem int

// Options for device operating system.
//
// Values of 500+ hold vendor-specific codes.
const (
	OSNotListed   OperatingSystem = 0  // Other Not Listed
	OS3DS         OperatingSystem = 1  // 3DS System Software
	OSAndroid     OperatingSystem = 2  // Android
	OSAppleTV     OperatingSystem = 3  // Apple TV Software
	OSAsha        OperatingSystem = 4  // Asha
	OSBada        OperatingSystem = 5  // Bada
	OSBlackBerry  OperatingSystem = 6  // BlackBerry
	OSBREW        OperatingSystem = 7  // BREW
	OSChromeOS    OperatingSystem = 8  // ChromeOS
	OSDarwin      OperatingSystem = 9  // Darwin
	OSFireOS      OperatingSystem = 10 // FireOS
	OSFirefoxOS   OperatingSystem = 11 // FirefoxOS
	OSHelenOS     OperatingSystem = 12 // HelenOS
	OSIOS         OperatingSystem = 13 // iOS
	OSLinux       OperatingSystem = 14 // Linux
	OSMacOS       OperatingSystem = 15 // MacOS
	OSMeeGo       OperatingSystem = 16 // MeeGo
	OSMorphOS     OperatingSystem = 17 // MorphOS
	OSNetBSD      OperatingSystem = 18 // NetBSD
	OSNucleusPLUS OperatingSystem = 19 // NucleusPLUS
	OSPSVita      OperatingSystem = 20 // PS Vita System Software
	OSPS3         OperatingSystem = 21 // PS3 System Software
	OSPS4         OperatingSystem = 22 // PS4 Software
	OSPSP         OperatingSystem = 23 // PSP System Software
	OSSymbian     OperatingSystem = 24 // Symbian
	OSTizen       OperatingSystem = 25 // Tizen
	OSWatchOS     OperatingSystem = 26 // WatchOS
	OSWebOS       OperatingSystem = 27 // WebOS
	OSWindows     OperatingSystem = 28 // Windows
)
