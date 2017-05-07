package openrtb_test

import (
	"encoding/json"

	"github.com/prebid/openrtb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RawJSON", func() {

	var _ json.Marshaler = (openrtb.RawJSON)(nil)
	var _ json.Unmarshaler = (*openrtb.RawJSON)(nil)

	It("should encode JSON", func() {
		subject := openrtb.RawJSON(`true`)

		actual, err := subject.MarshalJSON()
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal([]byte(`true`)))
	})

	It("should decode JSON", func() {
		subject := openrtb.RawJSON(nil)

		err := subject.UnmarshalJSON([]byte(`true`))
		Expect(err).NotTo(HaveOccurred())
		Expect(subject).To(Equal(openrtb.RawJSON(`true`)))
	})

	It("should decode JSON when embedded into struct", func() {
		wrapper := struct {
			Raw openrtb.RawJSON `json:"raw"`
		}{
			Raw: nil,
		}

		err := json.Unmarshal([]byte(`{"raw":true}`), &wrapper)
		Expect(err).NotTo(HaveOccurred())
		Expect(wrapper.Raw).To(Equal(openrtb.RawJSON(`true`)))
	})

	It("should encode JSON when embedded into struct", func() {
		wrapper := struct {
			Raw openrtb.RawJSON `json:"raw"`
		}{
			Raw: openrtb.RawJSON(`true`),
		}

		actual, err := json.Marshal(wrapper)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(MatchJSON(`{"raw":true}`))
	})

})
