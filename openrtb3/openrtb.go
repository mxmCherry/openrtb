// Package openrtb implements types for OpenRTB Specification v3.0
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/master/OpenRTB%20v3.0%20FINAL.md
package openrtb

// OpenRTB top-level object is the root for both request and response payloads.
// It includes versioning information and references to the Layer-4 domain model on which transactions are based.
// By default, the domain model used by OpenRTB is the Advertising Common Object Model (AdCOM).
//
// Some of these attributes are optional.
// The ver attribute, for example, indicates the OpenRTB specification version to which this payload conforms.
// This is also conveyed in Layer-1 via an HTTP header.
// Its utility here is more to assist in diagnostics by making the payload more self-documenting outside the context of a runtime transaction.
//
// The domainver attribute, however, does have runtime utility since the structures of Layer-4 objects may vary over time based on their specification versions.
// This attribute can assist in invoking the correct domain object parser or unmarshalling code.
type OpenRTB struct {
	// Attribute:
	//   ver
	// Type:
	//   string
	// Definition:
	//   Version of the Layer-3 OpenRTB specification (e.g., "3.0").
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   domainspec
	// Type:
	//   string; default “adcom”
	// Definition:
	//   Identifier of the Layer-4 domain model used to define items for sale, media associated with bids, etc.
	DomainSpec string `json:"domainspec,omitempty"`

	// Attribute:
	//   domainver
	// Type:
	//   string; required
	// Definition:
	//   Specification version of the Layer-4 domain model referenced in the domainspec attribute.
	DomainVer string `json:"domainver"`

	// Attribute:
	//   request
	// Type:
	//   object; required *
	// Definition:
	//   Bid request container.
	//   * Required only for request payloads.
	//   Refer to Object: Request.
	Request *Request `json:"request,omitempty"`

	// Attribute:
	//   response
	// Type:
	//   object; required *
	// Definition:
	//   Bid response container.
	//   * Required only for response payloads.
	//   Refer to Object: Response.
	Response *Response `json:"response,omitempty"`
}
