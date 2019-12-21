package sharing

type ArticleSharingService struct {
	sharer ArticleSharingRepo
}

func NewArticleSharingService(sharer ArticleSharingRepo) *ArticleSharingService {
	return &ArticleSharingService{sharer: sharer}
}

func (a ArticleSharingService) GetPublicEntries(page, limit int, sortBy string, order SortOrder) ([]PublicEntry, error) {
	offset := page * limit
	publicEntries, err := a.sharer.FetchPublicEntries(offset, limit, sortBy, order)
	if err != nil {
		return []PublicEntry{}, err
	}
	for _, entry := range publicEntries {
		art, err := a.sharer.FetchArticle(entry.ArticleID)
		comments, err := a.sharer.FetchComments(entry.EntryID, 0, 0)
		if err != nil {
			return []PublicEntry{}, err
		}
		entry.ContainedArticle = &art
		entry.Comments = &comments
	}
	return publicEntries, err
}
