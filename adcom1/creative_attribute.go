package adcom1

// CreativeAttribute specifies a standard list of creative attributes that can describe an actual ad or restrictions relative to a given placement.
type CreativeAttribute int

// Standard list of creative attributes that can describe an actual ad or restrictions relative to a given placement.
//
// Values of 500+ hold vendor-specific codes.
const (
	AttrAudioAuto              CreativeAttribute = 1  // Audio Ad (Autoplay)
	AttrAudioUser              CreativeAttribute = 2  // Audio Ad (User Initiated)
	AttrExpandableAuto         CreativeAttribute = 3  // Expandable (Automatic)
	AttrExpandableUserClick    CreativeAttribute = 4  // Expandable (User Initiated - Click)
	AttrExpandableUserRollover CreativeAttribute = 5  // Expandable (User Initiated - Rollover)
	AttrVideoAuto              CreativeAttribute = 6  // In-Banner Video Ad (Autoplay)
	AttrVideoUser              CreativeAttribute = 7  // In-Banner Video Ad (User Initiated)
	AttrPop                    CreativeAttribute = 8  // Pop (e.g., Over, Under, or Upon Exit)
	AttrProvocative            CreativeAttribute = 9  // Provocative or Suggestive Imagery
	AttrExtremeAnimation       CreativeAttribute = 10 // Shaky, Flashing, Flickering, Extreme Animation, Smileys
	AttrSurvey                 CreativeAttribute = 11 // Surveys
	AttrTextOnly               CreativeAttribute = 12 // Text Only
	AttrInteractive            CreativeAttribute = 13 // User Interactive (e.g., Embedded Games)
	AttrWindowsDialog          CreativeAttribute = 14 // Windows Dialog or Alert Style
	AttrHasAudioToggleButton   CreativeAttribute = 15 // Has Audio On/Off Button
	AttrHasSkipButton          CreativeAttribute = 16 // Ad Provides Skip Button (e.g. VPAID-rendered skip button on pre-roll video)
	AttrFlash                  CreativeAttribute = 17 // Adobe Flash
	AttrResponsive             CreativeAttribute = 18 // Responsive; Sizeless; Fluid (i.e., creatives that dynamically resize to environment)
)
