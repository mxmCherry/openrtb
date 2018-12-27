package adcom

// DistributionChannel is an abstraction of the various types of entities or channels through which ads are distributed.
// Examples include websites, mobile apps, and digital out-of-home (DOOH) systems.
// This generalized class contains those attributes and relationships that are common to all distribution channels and as such, all of its attributes and relationships are inherited by each of its derived classes.
//
// Note: As an abstract class, a DistributionChannel is never instantiated on its own.
// Only objects of its derived classes are actually realized.
type DistributionChannel struct {
	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Definition:
	//   Vendor-specific unique identifier of the distribution channel.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Definition:
	//   Displayable name of the distribution channel.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   pub
	// Type:
	//   object
	// Definition:
	//   Details about the publisher of the distribution channel.
	//   Refer to Object: Publisher.
	Pub *Publisher `json:"pub,omitempty"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Definition:
	//   Details about the content within the distribution channel.
	//   Refer to Object: Content.
	Content *Content `json:"content,omitempty"`
}
