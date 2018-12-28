# native1/request/testdata/v1.2

JSON examples copied from [OpenRTB Dynamic Native Ads API](https://iabtechlab.com/standards/openrtb-native/) [1.2](https://iabtechlab.com/wp-content/uploads/2016/07/OpenRTB-Native-Ads-Specification-Final-1.2.pdf) section "6 Bid Request/Response Samples".

Some empty/zero attributes were omited (not copied) because of [encoding/json](https://golang.org/pkg/encoding/json/) `omitempty` and [gomega.MatchJSON(...)](http://onsi.github.io/gomega/#matchjsonjson-interface).

Some parts were adapted due too poor quality of examples in spec:
- `eventtrackers`: called `eventtrackers` in spec, but `eventrackers` in example
- `eventtrackers`: request example included response `eventrrackers`
