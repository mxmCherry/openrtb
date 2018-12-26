package adcom

// ContentContext represents options for indicating the type of content being used or consumed by the user in which ads may appear.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type ContentContext int

// Options for indicating the type of content being used or consumed by the user in which ads may appear.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
const (
	ContentVideo   ContentContext = 1 // 1	Video (i.e., video file or stream such as Internet TV broadcasts)
	ContentGame    ContentContext = 2 // 2	Game (i.e., an interactive software game)
	ContentMusic   ContentContext = 3 // 3	Music (i.e., audio file or stream such as Internet radio broadcasts)
	ContentApp     ContentContext = 4 // 4	Application (i.e., an interactive software application)
	ContentText    ContentContext = 5 // 5	Text (i.e., primarily textual document such as a web page, eBook, or news article)
	ContentOther   ContentContext = 6 // 6	Other (i.e., none of the other categories applies)
	ContentUnknown ContentContext = 7 // 7	Unknown
)
