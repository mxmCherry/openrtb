package adcom1_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/mxmCherry/openrtb/v15/adcom1"
)

var _ = DescribeTable(
	"Marshaling",
	func(obj interface{}, file string) {
		golden, err := ioutil.ReadFile(filepath.Join("testdata", file))
		Expect(err).NotTo(HaveOccurred())

		Expect(json.Unmarshal(golden, obj)).To(Succeed())

		b, err := json.Marshal(obj)
		Expect(err).NotTo(HaveOccurred())

		Expect(b).To(MatchJSON(golden))
	},

	Entry("request context", new(RequestContext), "request-context.json"),
	Entry("item specifications", new(ItemSpec), "item-specifications.json"),
	Entry("media response", new(BidMedia), "media-response.json"),
)
