# testdata

JSON examples copied from [OpenRTB](https://www.iab.com/guidelines/real-time-bidding-rtb-project/) [v2.5](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf) spec - section 6. Bid Request/Response Samples.

Some empty/zero attributes were omited (not copied) because of [encoding/json](https://golang.org/pkg/encoding/json/) `omitempty` and [gomega.MatchJSON(...)](http://onsi.github.io/gomega/#matchjsonjson-interface).
