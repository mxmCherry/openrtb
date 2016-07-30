# Testdata

JSON examples copied from [OpenRTB](//github.com/openrtb/OpenRTB) [v2.3.1](//github.com/openrtb/OpenRTB/blob/master/OpenRTB-API-Specification-Version-2-3-1-FINAL.pdf) spec - section 6. Bid Request/Response Samples.

Some empty/zero attributes (like `imp[i].banner.pos == 0`) were omited because of [encoding/json](//golang.org/pkg/encoding/json/) `omitempty` and [gomega.MatchJSON(...)](//onsi.github.io/gomega/#matchjsonjson-interface).
