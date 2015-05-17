package rtb

// 5.8 Video Bid Response Protocols
//
// The following table lists the options for video bid response protocols that could be supported by an
// exchange.
const (
	VideoBidResponseProtocolVAST10        uint8 = 1 // 1 VAST 1.0
	VideoBidResponseProtocolVAST20        uint8 = 2 // 2 VAST 2.0
	VideoBidResponseProtocolVAST30        uint8 = 3 // 3 VAST 3.0
	VideoBidResponseProtocolVAST10Wrapper uint8 = 4 // 1 VAST 1.0 Wrapper
	VideoBidResponseProtocolVAST20Wrapper uint8 = 5 // 2 VAST 2.0 Wrapper
	VideoBidResponseProtocolVAST30Wrapper uint8 = 6 // 3 VAST 3.0 Wrapper
)
