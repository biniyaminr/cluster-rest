package sharing

type ArticleSharer interface {
}

type ArticleSharingRepo interface {
	FetchComments(entryID string) ([]Comment, error)
}
