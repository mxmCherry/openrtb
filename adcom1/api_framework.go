package adcom1

// APIFramework represents API frameworks either supported by a placement or required by an ad.
type APIFramework int

// API frameworks either supported by a placement or required by an ad.
//
// Values of 500+ hold vendor-specific codes.
const (
	APIVPAID1  APIFramework = 1 // VPAID 1.0
	APIVPAID2  APIFramework = 2 // VPAID 2.0
	APIMRAID1  APIFramework = 3 // MRAID 1.0
	APIORMMA   APIFramework = 4 // ORMMA
	APIMRAID2  APIFramework = 5 // MRAID 2.0
	APIMRAID3  APIFramework = 6 // MRAID 3.0
	APIOMID1   APIFramework = 7 // OMID 1.0
	APISIMID10 APIFramework = 8 // SIMID 1.0
	APISIMID11 APIFramework = 9 // SIMID 1.1
)
