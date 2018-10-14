package openrtb3

// NoBidReasonCode defines the options for a bidder to signal the exchange as to why it did not offer a bid for the item.
type NoBidReasonCode int64

// The following NoBidReasonCode constants list the options for a bidder to signal the exchange as to why it did not offer a bid for the item.
const (
	NoBidUnknownError                     NoBidReasonCode = 0  // Unknown Error
	NoBidTechnicalError                   NoBidReasonCode = 1  // Technical Error
	NoBidInvalidRequest                   NoBidReasonCode = 2  // Invalid Request
	NoBidKnownWebCrawler                  NoBidReasonCode = 3  // Known Web Crawler
	NoBidSuspectedNonHumanTraffic         NoBidReasonCode = 4  // Suspected Non-Human Traffic
	NoBidCloudDataCenterProxyIP           NoBidReasonCode = 5  // Cloud, Data Center, or Proxy IP
	NoBidUnsupportedDevice                NoBidReasonCode = 6  // Unsupported Device
	NoBidBlockedPublisherSite             NoBidReasonCode = 7  // Blocked Publisher or Site
	NoBidUnmatchedUser                    NoBidReasonCode = 8  // Unmatched User
	NoBidDailyUserCapMet                  NoBidReasonCode = 9  // Daily User Cap Met
	NoBidDailyDomainCapMet                NoBidReasonCode = 10 // Daily Domain Cap Met
	NoBidAdsTxtAuthorizationUnavailable   NoBidReasonCode = 11 // Ads.txt Authorization Unavailable
	NoBidAdsTxtAuthorizationViolation     NoBidReasonCode = 12 // Ads.txt Authorization Violation
	NoBidAdsCertAuthenticationUnavailable NoBidReasonCode = 13 // Ads.cert Authentication Unavailable
	NoBidAdsCertAuthenticationViolation   NoBidReasonCode = 14 // Ads.cert Authentication Violation
	NoBidInsufficientAuctionTime          NoBidReasonCode = 15 // Insufficient Auction Time

	// 500+ Exchange specific values; should be communicated with buyers a priori.
)
