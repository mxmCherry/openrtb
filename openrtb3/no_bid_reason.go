package openrtb3

// NoBidReason lists the options for a bidder to signal the exchange as to why it did not offer a bid for the item.
type NoBidReason int64

// NoBidReason options.
//
// Values 500+ are exchange specific values; should be communicated with buyers beforehand.
const (
	NoBidUnknownError              NoBidReason = 0  // Unknown Error
	NoBidTechnicalError            NoBidReason = 1  // Technical Error
	NoBidInvalidRequest            NoBidReason = 2  // Invalid Request
	NoBidCrawler                   NoBidReason = 3  // Known Web Crawler
	NoBidNonHuman                  NoBidReason = 4  // Suspected Non-Human Traffic
	NoBidProxy                     NoBidReason = 5  // Cloud, Data Center, or Proxy IP
	NoBidUnsupportedDevice         NoBidReason = 6  // Unsupported Device
	NoBidBlockedPublisher          NoBidReason = 7  // Blocked Publisher or Site
	NoBidUnmatchedUser             NoBidReason = 8  // Unmatched User
	NoBidDailyUserCap              NoBidReason = 9  // Daily User Cap Met
	NoBidDailyDomainCap            NoBidReason = 10 // Daily Domain Cap Met
	NoBidAuthorizationUnavailable  NoBidReason = 11 // Ads.txt Authorization Unavailable
	NoBidAuthorizationViolation    NoBidReason = 12 // Ads.txt Authorization Violation
	NoBidAuthenticationUnavailable NoBidReason = 13 // Ads.cert Authentication Unavailable
	NoBidAuthenticationViolation   NoBidReason = 14 // Ads.cert Authentication Violation
	NoBidInsufficientTime          NoBidReason = 15 // Insufficient Auction Time
	NoBidIncompleteSupplyChain     NoBidReason = 16 // Incomplete SupplyChain
	NoBidBlockedSupplyChainNode    NoBidReason = 17 // Blocked SupplyChain Node
)

// Ptr returns pointer to own value.
func (n NoBidReason) Ptr() *NoBidReason {
	return &n
}

// Val safely dereferences pointer, returning default value (NoBidUnknownError) for nil.
func (n *NoBidReason) Val() NoBidReason {
	if n == nil {
		return NoBidUnknownError
	}
	return *n
}
