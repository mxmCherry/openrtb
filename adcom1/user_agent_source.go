package adcom1

// UserAgentSource represents the source of a user agent.
type UserAgentSource int8

// Options for the user agent source.
const (
	UASourceUnknown     UserAgentSource = 0 // Unspecified/unknown
	UASourceLowEntropy  UserAgentSource = 1 // User-Agent Client Hints (only low-entropy headers were available)
	UASourceHighEntropy UserAgentSource = 2 // User-Agent Client Hints (with high-entropy headers available)
	UASourceParsed      UserAgentSource = 3 // Parsed from User-Agent header (the same string carried by the ua field)
)
