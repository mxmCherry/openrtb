package openrtb3

// LossReasonCode defines the options for an exchange to inform a bidder as to the reason why they did not win an item.
type LossReasonCode int64

// The following LossReasonCode constants list the options for an exchange to inform a bidder as to the reason why they did not win an item.
const (
	LossBidWon                                      LossReasonCode = 0   // Bid Won
	LossInternalError                               LossReasonCode = 1   // Internal Error
	LossImpressionOpportunityExpired                LossReasonCode = 2   // Impression Opportunity Expired
	LossInvalidBidResponse                          LossReasonCode = 3   // Invalid Bid Response
	LossInvalidDealID                               LossReasonCode = 4   // Invalid Deal ID
	LossInvalidAuctionID                            LossReasonCode = 5   // Invalid Auction ID
	LossInvalidAdvertiserDomain                     LossReasonCode = 6   // Invalid Advertiser Domain
	LossMissingMarkup                               LossReasonCode = 7   // Missing Markup
	LossMissingCreativeID                           LossReasonCode = 8   // Missing Creative ID
	LossMissingBidPrice                             LossReasonCode = 9   // Missing Bid Price
	LossMissingMininumCreativeApprovalData          LossReasonCode = 10  // Missing Minimum Creative Approval Data
	LossBidBelowAuctionFloor                        LossReasonCode = 100 // Bid was Below Auction Floor
	LossBidBelowDealFloor                           LossReasonCode = 101 // Bid was Below Deal Floor
	LossLostHigherBid                               LossReasonCode = 102 // Lost to Higher Bid
	LossLostBidDeal                                 LossReasonCode = 103 // Lost to a Bid for a Deal
	LossBuyerSeatBlocked                            LossReasonCode = 104 // Buyer Seat Blocked
	LossCreativeFilteredGeneralUnknown              LossReasonCode = 200 // Creative Filtered - General; Reason Unknown
	LossCreativeFilteredPendingProcessingExchange   LossReasonCode = 201 // Creative Filtered - Pending Processing by Exchange (e.g., approval, transcoding, etc.)
	LossCreativeFilteredDisapprovedExchange         LossReasonCode = 202 // Creative Filtered - Disapproved by Exchange
	LossCreativeFilteredSizeNotAllowed              LossReasonCode = 203 // Creative Filtered - Size Not Allowed
	LossCreativeFilteredIncorrectCreativeFormat     LossReasonCode = 204 // Creative Filtered - Incorrect Creative Format
	LossCreativeFilteredAdvertiserExclusions        LossReasonCode = 205 // Creative Filtered - Advertiser Exclusions
	LossCreativeFilteredNotSecure                   LossReasonCode = 206 // Creative Filtered - Not Secure
	LossCreativeFilteredLanguageExclusions          LossReasonCode = 207 // Creative Filtered - Language Exclusions
	LossCreativeFilteredCategoryExclusions          LossReasonCode = 208 // Creative Filtered - Category Exclusions
	LossCreativeFilteredCreativeAttributeExclusions LossReasonCode = 209 // Creative Filtered - Creative Attribute Exclusions
	LossCreativeFilteredAdTypeExclusions            LossReasonCode = 210 // Creative Filtered - Ad Type Exclusions
	LossCreativeFilteredAnimationTooLong            LossReasonCode = 211 // Creative Filtered - Animation Too Long
	LossCreativeFilteredNotAllowedDeal              LossReasonCode = 212 // Creative Filtered - Not Allowed in Deal

	// 500+	Exchange specific values; should be communicated with buyers a priori.
)
