package listing

//ArticleContents defines the basic contents of an article.
type ArticleContents struct {
	Title        string
	ArticleImage string
	Authors      []string
	Description  string
	TextContent  string
}
