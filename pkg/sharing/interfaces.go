package sharing

type ArticleSharer interface {
}

type ArticleSharingRepo interface {
	FetchComments(entryID string, offset, limit int) ([]Comment, error)
	FetchPublicEntries(offset, limit int, sortBy string, order SortOrder) ([]PublicEntry, error)
	FetchArticle(articleID int) (Article, error)
}
