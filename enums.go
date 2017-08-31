package openrtb

//enums and constants for tables 5.x.x (except 5.1 - IAB categories)

//5.2 Banner Ad Types

const (
	eBannerAdTypeXHTMLTextAd int8 = iota + 1
	eBannerAdTypeXHTMLBannerAd
	eBannerAdTypeJavaScriptAd
	eBannerAdTypeIframe
)

//5.3 Creative Attributes

const (
	eCreativeAttributesAudioAdAutoPlay int8 = iota + 1
	eCreativeAttributesAudioAdUserInitiated
	eCreativeAttributesExpandableAutomatic
	eCreativeAttributesExpandableClick
	eCreativeAttributesExpandableRollover
	eCreativeAttributesInBannerVideoAdAutoPlay
	eCreativeAttributesInBannerVideoAdUserInitiated
	eCreativeAttributesPop
	eCreativeAttributesProvocative
	eCreativeAttributesShaky
	eCreativeAttributesSurveys
	eCreativeAttributesTextOnly
	eCreativeAttributesUserInteractive
	eCreativeAttributesWindowsDialog
	eCreativeAttributesHasAudioOnOffButton
	eCreativeAttributesAdProvidesSkipButton
	eCreativeAttributesAdobeFlash
)

//5.4 Ad Position

const (
	eAdPositionUnknown int8 = iota
	eAdPositionAboveTheFold
	eAdPositionDepricated
	eAdPositionBelowTheFold
	eAdPositionHeader
	eAdPositionFooter
	eAdPositionSidebar
	eAdPositionFullScreen
)

//5.5 Expandable Direction

const (
	eExpandableDirectionLeft int8 = iota + 1
	eExpandableDirectionRight
	eExpandableDirectionUp
	eExpandableDirectionDown
	eExpandableDirectionFullScreen
)

//5.6 API Frameworks

const (
	eAPIFrameworksVPAID1 int8 = iota + 1
	eAPIFrameworksVPAID2
	eAPIFrameworksMRAID1
	eAPIFrameworksORMMA
	eAPIFrameworksMRAID2
	eAPIFrameworksMRAID3
)

//5.7 Video Linearity

const (
	eVideoLinearityLinear int8 = iota + 1
	eVideoLinearityNonLinear
)

//5.8 Protocols

const (
	eProtocolsVAST1 int8 = iota + 1
	eProtocolsVAST2
	eProtocolsVAST3
	eProtocolsVAST1Wrapper
	eProtocolsVAST2Wrapper
	eProtocolsVAST3Wrapper
	eProtocolsVAST4
	eProtocolsVAST4Wrapper
	eProtocolsDAAST1
	eProtocolsDAAST1Wrapper
)

//5.9 Video Placement Types

const (
	eVideoPlacementTypesInStream int8 = iota + 1
	eVideoPlacementTypesInBanner
	eVideoPlacementTypesInArticle
	eVideoPlacementTypesInFeed
	eVideoPlacementTypesInterstitialSliderFloating
)

//5.10 Playback Methods

const (
	ePlaybackMethodsInitiatesOnPageLoadWithSoundOn int8 = iota + 1
	ePlaybackMethodsInitiatesOnPageLoadWithSoundOff
	ePlaybackMethodsInitiatesOnClickWithSoundOn
	ePlaybackMethodsInitiatesOnMouseOverWithSoundOn
	ePlaybackMethodsInitiatesOnEnteringViewportWithSoundOn
	ePlaybackMethodsInitiatesOnEnteringViewportWithSoundOff
)

//5.11 Playback Cessation Modes

const (
	ePlaybackCessationModesOnVideoCompletion int8 = iota + 1
	ePlaybackCessationModesOnLeavingViewport
	ePlaybackCessationModesOnLeavingViewportContinuesAsFloatingSlider
)

//5.12 Start Delay

const (
	eStartDelayGenericPostRoll int8 = iota - 2
	eStartDelayGenericMidRoll
	eStartDelayPreRoll
)

//5.13 Production Quality

const (
	eProductionQualityUnknown int8 = iota
	eProductionQualityProfessionallyProduced
	eProductionQualityProsumer
	eProductionQualityUserGenerated
)

//5.14 Companion Types

const (
	eCompanionTypesStaticResource int8 = iota + 1
	eCompanionTypesHTMLResource
	eCompanionTypesIframeResource
)

//5.15 Content Delivery Methods

const (
	eContentDeliveryMethodsStreaming int8 = iota + 1
	eContentDeliveryMethodsProgressive
	eContentDeliveryMethodsDownload
)

//5.16 Feed Types

const (
	eFeedTypesMusicService int8 = iota + 1
	eFeedTypesFMAMBroadcast
	eFeedTypesPodcast
)

//5.17 Volume Normalization Modes

const (
	eVolumeNormalizationModesNone int8 = iota
	eVolumeNormalizationModesAdVolumeAverageNormalizedToContent
	eVolumeNormalizationModesAdVolumePeakNormalizedToContent
	eVolumeNormalizationModesAdLoudnessNormalizedToContent
	eVolumeNormalizationModesCustomVolumeNormalization
)

//5.18 Content Context

const (
	eContentContextVideo int8 = iota + 1
	eContentContextGame
	eContentContextMusic
	eContentContextApplication
	eContentContextText
	eContentContextOther
	eContentContextUnknown
)

//5.19 IQG Media Ratings

const (
	eIQGMediaRatingsAllAudiences int8 = iota + 1
	eIQGMediaRatingsEveryoneOver12
	eIQGMediaRatingsMatureAudiences
)

//5.20 Location Type

const (
	eLocationTypeGPS int8 = iota + 1
	eLocationTypeIP
	eLocationTypeUser
)

//5.21 Device Type

const (
	eDeviceTypeMobileTablet int8 = iota + 1
	eDeviceTypePersonalComputer
	eDeviceTypeConnectedTV
	eDeviceTypePhone
	eDeviceTypeTablet
	eDeviceTypeConnectedDevice
	eDeviceTypeSetTopBox
)

//5.22 Connection Type

const (
	eConnectionTypeUnknown int8 = iota
	eConnectionTypeEthernet
	eConnectionTypeWIFI
	eConnectionTypeCellularNetworkUnknownGeneration
	eConnectionTypeCellularNetwork2G
	eConnectionTypeCellularNetwork3G
	eConnectionTypeCellularNetwork4G
)

//5.23 IP Location Services

const (
	eIPLocationServiceIP2Location int8 = iota + 1
	eIPLocationServiceNeustar
	IPLocationServiceMaxMind
	IPLocationServiceNetAcuity
)

//5.24 No-Bid Reason Codes

const (
	eNoBidReasonUnknownError int8 = iota
	eNoBidReasonCodeTechnicalError
	eNoBidReasonCodeInvalidRequest
	eNoBidReasonCodeKnownWebSpider
	eNoBidReasonCodeSuspectedNonHumanTraffic
	eNoBidReasonCodeCloudDataCenterProxyIP
	eNoBidReasonCodeUnsupportedDevice
	eNoBidReasonCodeBlockedPublisherOrSite
	eNoBidReasonCodeUnmatchedUser
	eNoBidReasonCodeDailyReaderCapMet
	eNoBidReasonCodeDailyDomainCapMet
)

//5.25 Loss Reason Codes

const (
	eLossReasonCodesBidWon int8 = iota
	eLossReasonCodesInternalError
	eLossReasonCodesImpressionOpportunityExpired
	eLossReasonCodesInvalidBidResponse
	eLossReasonCodesInvalidDealID
	eLossReasonCodesInvalidAuctionID
	eLossReasonCodesInvalidAdvertiserDomain
	eLossReasonCodesMissingMarkup
	eLossReasonCodesMissingCreativeID
	eLossReasonCodesMissingBidPrice
	eLossReasonCodesMissingMinimumCreativeApprovalData
	eLossReasonCodesBidWasBelowAuctionFloor int8 = iota + 100
	eLossReasonCodesBidWasBelowDealFloor
	eLossReasonCodesLostToHigherBid
	eLossReasonCodesLostToABidForPMPDeal
	eLossReasonCodesBuyerSeatBlocked
	eLossReasonCodesCreativeFilteredGeneral uint8 = iota + 200
	eLossReasonCodesCreativeFilteredPendingProcessingByExchange
	eLossReasonCodesCreativeFilteredDisapprovedByExchange
	eLossReasonCodesCreativeFilteredSizeNotAllowed
	eLossReasonCodesCreativeFilteredIncorrectCreativeFormat
	eLossReasonCodesCreativeFilteredAdvertiserExclusions
	eLossReasonCodesCreativeFilteredAppBundleExclusions
	eLossReasonCodesCreativeFilteredNotSecure
	eLossReasonCodesCreativeFilteredLanguageExclusions
	eLossReasonCodesCreativeFilteredCategoryExclusions
	eLossReasonCodesCreativeFilteredCreativeAttributeExclusions
	eLossReasonCodesCreativeFilteredAdTypeExclusions
	eLossReasonCodesCreativeFilteredAnimationTooLong
	eLossReasonCodesCreativeFilteredNotAllowedInPMPDeal
)
