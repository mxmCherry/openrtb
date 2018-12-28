package adcom1

// DisplayContextType represents types of context in which a native ad may appear (i.e., the type of content surrounding the ad on the page).
// This is intended to denote primary content although other content may also appear on the page.
// Note that there are two levels of detail grouped by 10s (i.e., 12 is a refined case of 100).
type DisplayContextType int

// Types of context in which a native ad may appear (i.e., the type of content surrounding the ad on the page).
//
// Values of 500+ hold vendor-specific codes.
const (
	DisplayContextContent              DisplayContextType = 10 // Content-centric context (e.g., newsfeed, article, image gallery, video gallery, etc.).
	DisplayContextContentArticle       DisplayContextType = 11 // - Primarily article content, which could include images, etc. as part of the article.
	DisplayContextContentVideo         DisplayContextType = 12 // - Primarily video content.
	DisplayContextContentAudio         DisplayContextType = 13 // - Primarily audio content.
	DisplayContextContentImage         DisplayContextType = 14 // - Primarily image content.
	DisplayContextContentUserGenerated DisplayContextType = 15 // - User-generated content (e.g., forums, comments, etc.).
	DisplayContextSocial               DisplayContextType = 20 // Social-centric context (e.g., social network feed, email, chat, etc.).
	DisplayContextSocialEmail          DisplayContextType = 21 // - Primarily email content.
	DisplayContextSocialChat           DisplayContextType = 22 // - Primarily chat/IM content.
	DisplayContextProduct              DisplayContextType = 30 // Product context (e.g., product listings, details, recommendations, reviews, etc.).
	DisplayContextProductApp           DisplayContextType = 31 // - App store/marketplace.
	DisplayContextProductReview        DisplayContextType = 32 // - Product reviews site primarily, which may sell product secondarily.
)
