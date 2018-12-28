package adcom1

// DOOHVenueType represents the digital out-of-home venue types and is derived from DPAA Programmatic Standards.
type DOOHVenueType int

// Digital out-of-home venue types.
//
// Values of 500+ hold vendor-specific codes.
const (
	VenueAirborne                DOOHVenueType = 1  // Airborne
	VenueAirportsGeneral         DOOHVenueType = 2  // Airports - General
	VenueAirportsBaggageClaim    DOOHVenueType = 3  // Airports - Baggage Claim
	VenueAirportsTerminal        DOOHVenueType = 4  // Airports - Terminal
	VenueAirportsLounge          DOOHVenueType = 5  // Airports - Lounges
	VenueATM                     DOOHVenueType = 6  // ATMs
	VenueBacklight               DOOHVenueType = 7  // Backlights
	VenueBars                    DOOHVenueType = 8  // Bars
	VenueBench                   DOOHVenueType = 9  // Benches
	VenueBikeRack                DOOHVenueType = 10 // Bike Racks
	VenueBulletin                DOOHVenueType = 11 // Bulletins
	VenueBuses                   DOOHVenueType = 12 // Buses
	VenueCafes                   DOOHVenueType = 13 // Cafes
	VenueCasualDining            DOOHVenueType = 14 // Casual Dining Restaurants
	VenueChildCare               DOOHVenueType = 15 // Child Care
	VenueCinema                  DOOHVenueType = 16 // Cinema
	VenueCityInformationPanel    DOOHVenueType = 17 // City Information Panels
	VenueConvenienceStore        DOOHVenueType = 18 // Convenience Stores
	VenueDedicatedWildPosting    DOOHVenueType = 19 // Dedicated Wild Posting
	VenueDoctorsOffice           DOOHVenueType = 20 // Doctors Offices - General
	VenueDoctorsOfficeObstetrics DOOHVenueType = 21 // Doctors Offices - Obstetrics
	VenueDoctorsOfficePediatrics DOOHVenueType = 22 // Doctors Offices - Pediatrics
	VenueFamilyEntertainment     DOOHVenueType = 23 // Family entertainment
	VenueFerry                   DOOHVenueType = 24 // Ferries
	VenueFinancialService        DOOHVenueType = 25 // Financial Services
	VenueGasStation              DOOHVenueType = 26 // Gas Stations
	VenueGolfCourse              DOOHVenueType = 27 // Golf Courses
	VenueGym                     DOOHVenueType = 28 // Gyms
	VenueHospital                DOOHVenueType = 29 // Hospitals
	VenueHotel                   DOOHVenueType = 30 // Hotels
	VenueJuniorPoster            DOOHVenueType = 31 // Junior Posters
	VenueKiosk                   DOOHVenueType = 32 // Kiosks
	VenueMall                    DOOHVenueType = 33 // Malls - General
	VenueMallFoodCourt           DOOHVenueType = 34 // Malls - Food Courts
	VenueMarine                  DOOHVenueType = 35 // Marine
	VenueMobileBillboard         DOOHVenueType = 36 // Mobile Billboards
	VenueMovieTheaterLobby       DOOHVenueType = 37 // Movie Theater Lobbies
	VenueNewsStand               DOOHVenueType = 38 // Newsstands
	VenueOfficeBuilding          DOOHVenueType = 39 // Office Buildings
	VenuePhoneKiosk              DOOHVenueType = 40 // Phone Kiosks
	VenuePoster                  DOOHVenueType = 41 // Posters
	VenueQSR                     DOOHVenueType = 42 // QSR [Quick-Service Restaurants / Fast-Food]
	VenueRail                    DOOHVenueType = 43 // Rail
	VenueReceptacle              DOOHVenueType = 44 // Receptacles
	VenueResortLeisure           DOOHVenueType = 45 // Resorts / Leisure
	VenueRetail                  DOOHVenueType = 46 // Retail
	VenueSalon                   DOOHVenueType = 47 // Salons
	VenueShelter                 DOOHVenueType = 48 // Shelters
	VenueSportsArena             DOOHVenueType = 49 // Sports Arenas
	VenueSubway                  DOOHVenueType = 50 // Subway
	VenueTaxi                    DOOHVenueType = 51 // Taxis / Wrapped vehicles
	VenueTruckSide               DOOHVenueType = 52 // Truckside
	VenueUniversity              DOOHVenueType = 53 // Universities
	VenueUrbanPanel              DOOHVenueType = 54 // Urban Panels
	VenueVeterinarianOffice      DOOHVenueType = 55 // Veterinarian Offices
	VenueWallSpectacular         DOOHVenueType = 56 // Walls / Spectaculars
	VenueOther                   DOOHVenueType = 57 // Other
)
