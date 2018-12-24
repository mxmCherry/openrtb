package openrtb

// LossReason represents OpenRTB Loss Reason Code enumeration.
//
// It lists the options for an exchange to inform a bidder as to the reason why they did not win an item.
type LossReason int64

// LossReason options.
//
// Values of 500+	are exchange specific values; should be communicated with buyers a priori.
const (
	LossWon                     LossReason = 0   // Bid Won
	LossInternalError           LossReason = 1   // Internal Error
	LossExpired                 LossReason = 2   // Impression Opportunity Expired
	LossInvalidResponse         LossReason = 3   // Invalid Bid Response
	LossInvalidDealID           LossReason = 4   // Invalid Deal ID
	LossInvalidAuctionID        LossReason = 5   // Invalid Auction ID
	LossInvalidAdvertiserDomain LossReason = 6   // Invalid Advertiser Domain
	LossMissingMarkup           LossReason = 7   // Missing Markup
	LossMissingCreativeID       LossReason = 8   // Missing Creative ID
	LossMissingBidPrice         LossReason = 9   // Missing Bid Price
	LossMissingApproval         LossReason = 10  // Missing Minimum Creative Approval Data
	LossBelowAuctionFloor       LossReason = 100 // Bid was Below Auction Floor
	LossBelowDealFloor          LossReason = 101 // Bid was Below Deal Floor
	LossLostToHigherBid         LossReason = 102 // Lost to Higher Bid
	LossLostToDealBid           LossReason = 103 // Lost to a Bid for a Deal
	LossSeatBlocked             LossReason = 104 // Buyer Seat Blocked
	LossCreativeFiltered        LossReason = 200 // Creative Filtered - General; Reason Unknown
	LossPendingProcessing       LossReason = 201 // Creative Filtered - Pending Processing by Exchange (e.g., approval, transcoding, etc.)
	LossDisapproved             LossReason = 202 // Creative Filtered - Disapproved by Exchange
	LossSizeNotAllowed          LossReason = 203 // Creative Filtered - Size Not Allowed
	LossIncorrectFormat         LossReason = 204 // Creative Filtered - Incorrect Creative Format
	LossAdvertiserExclusions    LossReason = 205 // Creative Filtered - Advertiser Exclusions
	LossNotSecure               LossReason = 206 // Creative Filtered - Not Secure
	LossLanguageExclusions      LossReason = 207 // Creative Filtered - Language Exclusions
	LossCategoryExclusions      LossReason = 208 // Creative Filtered - Category Exclusions
	LossAttributeExclusions     LossReason = 209 // Creative Filtered - Creative Attribute Exclusions
	LossAdTypeExclusions        LossReason = 210 // Creative Filtered - Ad Type Exclusions
	LossAnimationTooLong        LossReason = 211 // Creative Filtered - Animation Too Long
	LossNotAllowedInDeal        LossReason = 212 // Creative Filtered - Not Allowed in Deal
)
