package listing

//Feed struct represents a single feed with its contents
type Feed struct {
	Title    string
	Link     string
	Articles []Article
}

//FeedInformation represents detailed information about a feed
type FeedInformation struct {
	Title       string
	Link        string
	Description string
	ImageLink   string
	Language    string
}

//Article represents a single article
type Article struct {
	Title           string
	Link            string
	Description     string
	PublicationDate string
}
