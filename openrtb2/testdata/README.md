# testdata

JSON examples copied from [OpenRTB](https://iabtechlab.com/standards/openrtb/) spec documents:
- [2.5](https://iabtechlab.com/wp-content/uploads/2016/07/OpenRTB-API-Specification-Version-2-5-FINAL.pdf) spec - section 6. Bid Request/Response Samples
- [2.6](https://iabtechlab.com/wp-content/uploads/2022/04/OpenRTB-2-6_FINAL.pdf) spec - section 6. Bid Request/Response Samples

Some empty/zero attributes were omited (not copied) because of [encoding/json](https://golang.org/pkg/encoding/json/) `omitempty` and [gomega.MatchJSON(...)](http://onsi.github.io/gomega/#matchjsonjson-interface).

For [2.6](https://iabtechlab.com/wp-content/uploads/2022/04/OpenRTB-2-6_FINAL.pdf) the following adjustments were applied:

- [bid-request/2.6/simple-banner.json](bid-request/2.6/simple-banner.json) - spec example embedded `user` into `site`, moved `user` to top level
- [bid-request/2.6/expandable-creative.json](bid-request/2.6/expandable-creative.json) - `data[2].value` added, moved to `data[2].segment[0].value`
- [bid-request/2.6/mobile.json](bid-request/2.6/mobile.json) - `imp[0].instl = 0` removed (because of JSON/omitempty)
- [bid-request/2.6/VIDEO.json](bid-request/2.6/VIDEO.json) - `imp[0].video.apis` renamed to `api` (it is singular in spec, but plural in example)
