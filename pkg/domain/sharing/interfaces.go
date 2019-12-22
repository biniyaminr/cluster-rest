package sharing

type ArticleSharer interface {
	GetPublicEntries(page, limit int, sortBy string, order SortOrder) ([]PublicEntry, error)
	GetPublicEntry(entryID int) (PublicEntry, error)
	AddPublicEntry(sourceRssID int, sharedBy, title, link string) (PublicEntry, error)
	ShareToUser(senderUsername, receiverUsername, message string, articleID int) (Recommendation, error)
}

type ArticleSharingRepo interface {
	FetchComments(entryID int, offset, limit int) ([]Comment, error)
	FetchPublicEntries(offset, limit int, sortBy string, order SortOrder) ([]PublicEntry, error)
	FetchPublicEntry(entryID int) (PublicEntry, error)
	FetchArticle(articleID int) (Article, error)
	AddArticle(article *Article) (int, error)
	AddPublicEntry(entry *PublicEntry) (int, error)
	AddRecommendation(recommendation Recommendation) error
	AddLike(entryID int, username string) error
	DeleteLike(entryID, username string) error
}
