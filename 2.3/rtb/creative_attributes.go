package rtb

// 5.3 Creative Attributes
//
// The following table specifies a standard list of creative attributes that can describe an ad being served or
// serve as restrictions of thereof.
//
// TODO: review and rename
const (
	CreativeAttributeAudioAutoPlay                   uint8 = 1  // 1 Audio Ad (Auto-Play)
	CreativeAttributeAudioUserInitiated              uint8 = 2  // 2 Audio Ad (User Initiated)
	CreativeAttributeExpandableAutomatic             uint8 = 3  // 3 Expandable (Automatic)
	CreativeAttributeExpandableUserInitiatedClick    uint8 = 4  // 4 Expandable (User Initiated - Click)
	CreativeAttributeExpandableUserInitiatedRollover uint8 = 5  // 5 Expandable (User Initiated - Rollover)
	CreativeAttributeInBannerVideoAutoPlay           uint8 = 6  // 6 In-Banner Video Ad (Auto-Play)
	CreativeAttributeInBannerVideoUserInitiated      uint8 = 7  // 7 In-Banner Video Ad (User Initiated)
	CreativeAttributePop                             uint8 = 8  // 8 Pop (e.g., Over, Under, or Upon Exit)
	CreativeAttributeProvocative                     uint8 = 9  // 9 Provocative or Suggestive Imagery
	CreativeAttributeExtremeAnimation                uint8 = 10 // 10 Shaky, Flashing, Flickering, Extreme Animation, Smileys
	CreativeAttributeSurvey                          uint8 = 11 // 11 Surveys
	CreativeAttributeTextOnly                        uint8 = 12 // 12 Text Only
	CreativeAttributeUserInteractive                 uint8 = 13 // 13 User Interactive (e.g., Embedded Games)
	CreativeAttributeNativeStyle                     uint8 = 14 // 14 Windows Dialog or Alert Style
	CreativeAttributeHasAudioOnOffButton             uint8 = 15 // 15 Has Audio On/Off Button
	CreativeAttributeCanBeSkipped                    uint8 = 16 // 16 Ad Can be Skipped (e.g., Skip Button on Pre-Roll Video)
)
