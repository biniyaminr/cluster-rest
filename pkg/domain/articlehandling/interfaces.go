package articlehandling

type ArticleFetcherRepo interface {
	FetchArticleContents(url string) (ArticleContents, error)
	FetchMultipleArticlesContents(urls ...string) (map[string]ArticleContents, error)
	FetchArticleImage(url string) (string, error)
	FetchMultipleArticleImages(url ...string) (map[string]string, error)
}

type ArticleLister interface {
	FetchArticleContents(url string) (ArticleContents, error)
	FetchMultipleArticleContents(urls []string) (map[string]ArticleContents, error)
	FetchArticleImage(url string) (string, error)
	FetchMultipleArticleImages(articleURLs []string) (map[string]string, error)
}
