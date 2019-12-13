package repository

//ArticleContents represents the basic contents of a given article
type ArticleContents struct {
	Title        string
	ArticleImage string
	Authors      []string
	Description  string
	TextContent  string
}
