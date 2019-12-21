package sharing

type ArticleSharer interface {
	GetPublicEntries(page, limit int, sortBy string, order SortOrder) ([]PublicEntry, error)
}

type ArticleSharingRepo interface {
	FetchComments(entryID int, offset, limit int) ([]Comment, error)
	FetchPublicEntries(offset, limit int, sortBy string, order SortOrder) ([]PublicEntry, error)
	FetchArticle(articleID int) (Article, error)
}
