# v1.2

JSON examples copied from [OpenRTB](https://www.iab.com/guidelines/real-time-bidding-rtb-project/) [Dynamic Native Ads API
Specification Version 1.2](https://www.iab.com/wp-content/uploads/2018/03/OpenRTB-Native-Ads-Specification-Final-1.2.pdf) spec - section 6 Bid Request/Response Samples.

Some empty/zero attributes were omited (not copied) because of [encoding/json](https://golang.org/pkg/encoding/json/) `omitempty` and [gomega.MatchJSON(...)](http://onsi.github.io/gomega/#matchjsonjson-interface).

Some parts were adapted due too poor quality of examples in spec:
- `eventtrackers`: called `eventtrackers` in spec, but `eventrackers` in example
- `eventtrackers`: request example included response `eventrrackers`
