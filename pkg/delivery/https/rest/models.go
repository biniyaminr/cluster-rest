package rest

type Response struct {
	Status  string            `json:"status"`
	Data    []ArticleContents `json:"data"`
	Message string            `json:"message"`
}

type ArticleContents struct {
	Title        string   `json:"title"`
	ArticleImage string   `json:"article_image"`
	Authors      []string `json:"authors"`
	Description  string   `json:"description"`
	TextContent  string   `json:"text_content"`
}
