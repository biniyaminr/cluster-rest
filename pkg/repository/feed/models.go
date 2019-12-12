package repository

//MappedFeed struct represents a single feed along with the xml mapping
type MappedFeed struct {
	Title    string          `xml:"channel>title"`
	Link     string          `xml:"channel>link"`
	Articles []MappedArticle `xml:"channel>item"`
}

type MappedFeedInformation struct {
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	ImageLink   string `xml:"channel>image>url"`
	Language    string `xml:"channel>language"`
}

//MappedArticle represents a single article along with the xml mapping
type MappedArticle struct {
	Title           string `xml:"title"`
	Link            string `xml:"link"`
	Description     string `xml:"description"`
	PublicationDate string `xml:"pubDate"`
}
