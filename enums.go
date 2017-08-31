package openrtb

//enums and constants for tables 5.x.x (except 5.1 - IAB categories)

//5.2 Banner Ad Types

const (
	EBannerAdTypeXHTMLTextAd int8 = iota + 1
	EBannerAdTypeXHTMLBannerAd
	EBannerAdTypeJavaScriptAd
	EBannerAdTypeIframe
)

//5.3 Creative Attributes

const (
	ECreativeAttributesAudioAdAutoPlay int8 = iota + 1
	ECreativeAttributesAudioAdUserInitiated
	ECreativeAttributesExpandableAutomatic
	ECreativeAttributesExpandableClick
	ECreativeAttributesExpandableRollover
	ECreativeAttributesInBannerVideoAdAutoPlay
	ECreativeAttributesInBannerVideoAdUserInitiated
	ECreativeAttributesPop
	ECreativeAttributesProvocative
	ECreativeAttributesShaky
	ECreativeAttributesSurveys
	ECreativeAttributesTextOnly
	ECreativeAttributesUserInteractive
	ECreativeAttributesWindowsDialog
	ECreativeAttributesHasAudioOnOffButton
	ECreativeAttributesAdProvidesSkipButton
	ECreativeAttributesAdobeFlash
)

//5.4 Ad Position

const (
	EAdPositionUnknown int8 = iota
	EAdPositionAboveTheFold
	EAdPositionDepricated
	EAdPositionBelowTheFold
	EAdPositionHeader
	EAdPositionFooter
	EAdPositionSidebar
	EAdPositionFullScreen
)

//5.5 Expandable Direction

const (
	EExpandableDirectionLeft int8 = iota + 1
	EExpandableDirectionRight
	EExpandableDirectionUp
	EExpandableDirectionDown
	EExpandableDirectionFullScreen
)

//5.6 API Frameworks

const (
	EAPIFrameworksVPAID1 int8 = iota + 1
	EAPIFrameworksVPAID2
	EAPIFrameworksMRAID1
	EAPIFrameworksORMMA
	EAPIFrameworksMRAID2
	EAPIFrameworksMRAID3
)

//5.7 Video Linearity

const (
	EVideoLinearityLinear int8 = iota + 1
	EVideoLinearityNonLinear
)

//5.8 Protocols

const (
	EProtocolsVAST1 int8 = iota + 1
	EProtocolsVAST2
	EProtocolsVAST3
	EProtocolsVAST1Wrapper
	EProtocolsVAST2Wrapper
	EProtocolsVAST3Wrapper
	EProtocolsVAST4
	EProtocolsVAST4Wrapper
	EProtocolsDAAST1
	EProtocolsDAAST1Wrapper
)

//5.9 Video Placement Types

const (
	EVideoPlacementTypesInStream int8 = iota + 1
	EVideoPlacementTypesInBanner
	EVideoPlacementTypesInArticle
	EVideoPlacementTypesInFeed
	EVideoPlacementTypesInterstitialSliderFloating
)

//5.10 Playback Methods

const (
	EPlaybackMethodsInitiatesOnPageLoadWithSoundOn int8 = iota + 1
	EPlaybackMethodsInitiatesOnPageLoadWithSoundOff
	EPlaybackMethodsInitiatesOnClickWithSoundOn
	EPlaybackMethodsInitiatesOnMouseOverWithSoundOn
	EPlaybackMethodsInitiatesOnEnteringViewportWithSoundOn
	EPlaybackMethodsInitiatesOnEnteringViewportWithSoundOff
)

//5.11 Playback Cessation Modes

const (
	EPlaybackCessationModesOnVideoCompletion int8 = iota + 1
	EPlaybackCessationModesOnLeavingViewport
	EPlaybackCessationModesOnLeavingViewportContinuesAsFloatingSlider
)

//5.12 Start Delay

const (
	EStartDelayGenericPostRoll int8 = iota - 2
	EStartDelayGenericMidRoll
	EStartDelayPreRoll
)

//5.13 Production Quality

const (
	EProductionQualityUnknown int8 = iota
	EProductionQualityProfessionallyProduced
	EProductionQualityProsumer
	EProductionQualityUserGenerated
)

//5.14 Companion Types

const (
	ECompanionTypesStaticResource int8 = iota + 1
	ECompanionTypesHTMLResource
	ECompanionTypesIframeResource
)

//5.15 Content Delivery Methods

const (
	EContentDeliveryMethodsStreaming int8 = iota + 1
	EContentDeliveryMethodsProgressive
	EContentDeliveryMethodsDownload
)

//5.16 Feed Types

const (
	EFeedTypesMusicService int8 = iota + 1
	EFeedTypesFMAMBroadcast
	EFeedTypesPodcast
)

//5.17 Volume Normalization Modes

const (
	EVolumeNormalizationModesNone int8 = iota
	EVolumeNormalizationModesAdVolumeAverageNormalizedToContent
	EVolumeNormalizationModesAdVolumePeakNormalizedToContent
	EVolumeNormalizationModesAdLoudnessNormalizedToContent
	EVolumeNormalizationModesCustomVolumeNormalization
)

//5.18 Content Context

const (
	EContentContextVideo int8 = iota + 1
	EContentContextGame
	EContentContextMusic
	EContentContextApplication
	EContentContextText
	EContentContextOther
	EContentContextUnknown
)

//5.19 IQG Media Ratings

const (
	EIQGMediaRatingsAllAudiences int8 = iota + 1
	EIQGMediaRatingsEveryoneOver12
	EIQGMediaRatingsMatureAudiences
)

//5.20 Location Type

const (
	ELocationTypeGPS int8 = iota + 1
	ELocationTypeIP
	ELocationTypeUser
)

//5.21 Device Type

const (
	EDeviceTypeMobileTablet int8 = iota + 1
	EDeviceTypePersonalComputer
	EDeviceTypeConnectedTV
	EDeviceTypePhone
	EDeviceTypeTablet
	EDeviceTypeConnectedDevice
	EDeviceTypeSetTopBox
)

//5.22 Connection Type

const (
	EConnectionTypeUnknown int8 = iota
	EConnectionTypeEthernet
	EConnectionTypeWIFI
	EConnectionTypeCellularNetworkUnknownGeneration
	EConnectionTypeCellularNetwork2G
	EConnectionTypeCellularNetwork3G
	EConnectionTypeCellularNetwork4G
)

//5.23 IP Location Services

const (
	EIPLocationServiceIP2Location int8 = iota + 1
	EIPLocationServiceNeustar
	IPLocationServiceMaxMind
	IPLocationServiceNetAcuity
)

//5.24 No-Bid Reason Codes

const (
	ENoBidReasonUnknownError int8 = iota
	ENoBidReasonCodeTechnicalError
	ENoBidReasonCodeInvalidRequest
	ENoBidReasonCodeKnownWebSpider
	ENoBidReasonCodeSuspectedNonHumanTraffic
	ENoBidReasonCodeCloudDataCenterProxyIP
	ENoBidReasonCodeUnsupportedDevice
	ENoBidReasonCodeBlockedPublisherOrSite
	ENoBidReasonCodeUnmatchedUser
	ENoBidReasonCodeDailyReaderCapMet
	ENoBidReasonCodeDailyDomainCapMet
)

//5.25 Loss Reason Codes

const (
	ELossReasonCodesBidWon int8 = iota
	ELossReasonCodesInternalError
	ELossReasonCodesImpressionOpportunityExpired
	ELossReasonCodesInvalidBidResponse
	ELossReasonCodesInvalidDealID
	ELossReasonCodesInvalidAuctionID
	ELossReasonCodesInvalidAdvertiserDomain
	ELossReasonCodesMissingMarkup
	ELossReasonCodesMissingCreativeID
	ELossReasonCodesMissingBidPrice
	ELossReasonCodesMissingMinimumCreativeApprovalData
	ELossReasonCodesBidWasBelowAuctionFloor int8 = iota + 100
	ELossReasonCodesBidWasBelowDealFloor
	ELossReasonCodesLostToHigherBid
	ELossReasonCodesLostToABidForPMPDeal
	ELossReasonCodesBuyerSeatBlocked
	ELossReasonCodesCreativeFilteredGeneral uint8 = iota + 200
	ELossReasonCodesCreativeFilteredPendingProcessingByExchange
	ELossReasonCodesCreativeFilteredDisapprovedByExchange
	ELossReasonCodesCreativeFilteredSizeNotAllowed
	ELossReasonCodesCreativeFilteredIncorrectCreativeFormat
	ELossReasonCodesCreativeFilteredAdvertiserExclusions
	ELossReasonCodesCreativeFilteredAppBundleExclusions
	ELossReasonCodesCreativeFilteredNotSecure
	ELossReasonCodesCreativeFilteredLanguageExclusions
	ELossReasonCodesCreativeFilteredCategoryExclusions
	ELossReasonCodesCreativeFilteredCreativeAttributeExclusions
	ELossReasonCodesCreativeFilteredAdTypeExclusions
	ELossReasonCodesCreativeFilteredAnimationTooLong
	ELossReasonCodesCreativeFilteredNotAllowedInPMPDeal
)
