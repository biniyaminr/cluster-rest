package articlehandling

//ArticleContents defines the basic contents of an article.
type ArticleContents struct {
	Title        string   `json:"title" xml:"title"`
	ArticleImage string   `json:"article_image" xml:"article_image"`
	Authors      []string `json:"authors" xml:"authors"`
	Description  string   `json:"description" xml:"description"`
	TextContent  string   `json:"text_content" xml:"text_content"`
}
