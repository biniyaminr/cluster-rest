package listing

//ArticleContents defines the basic contents of an article.
type ArticleContents struct {
	Title        string   `json:"title"`
	ArticleImage string   `json:"article_image"`
	Authors      []string `json:"authors"`
	Description  string   `json:"description"`
	TextContent  string   `json:"text_content"`
}
