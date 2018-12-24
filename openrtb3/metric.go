package openrtb

import "encoding/json"

// Metric object is associated with an item as an array of metrics.
// These metrics can offer insight to assist with decisioning such as average recent viewability, click-through rate, etc.
// Each metric is identified by its type, reports the value of the metric, and optionally identifies the source or vendor measuring the value.
type Metric struct {
	// Attribute:
	//   type
	// Type:
	//   string; required
	// Definition:
	//   Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Type string `json:"type"`

	// Attribute:
	//   value
	// Type:
	//   float; required
	// Definition:
	//   Number representing the value of the metric.
	//   Probabilities must be in the range 0.0 – 1.0.
	Value float64 `json:"value,omitempty"`

	// Attribute:
	//   vendor
	// Type:
	//   string; recommended
	// Definition:
	//   Source of the value using exchange curated string names which should be published to bidders a priori.
	//   If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	Vendor string `json:"vendor,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional exchange-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
