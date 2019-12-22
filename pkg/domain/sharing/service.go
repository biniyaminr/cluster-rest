package sharing

import "time"

type ArticleSharingService struct {
	sharingRepo ArticleSharingRepo
}

func NewArticleSharingService(sharer ArticleSharingRepo) *ArticleSharingService {
	return &ArticleSharingService{sharingRepo: sharer}
}

//GetPublicEntries returns an array of the specified number of public entries
func (a ArticleSharingService) GetPublicEntries(page, limit int, sortBy string, order SortOrder) ([]PublicEntry, error) {
	offset := page * limit
	publicEntries, err := a.sharingRepo.FetchPublicEntries(offset, limit, sortBy, order)
	if err != nil {
		return []PublicEntry{}, err
	}
	for _, entry := range publicEntries {
		art, err := a.sharingRepo.FetchArticle(entry.ArticleID)
		comments, err := a.sharingRepo.FetchComments(entry.EntryID, 0, 0)
		if err != nil {
			return []PublicEntry{}, err
		}
		entry.ContainedArticle = &art
		entry.Comments = &comments
	}
	return publicEntries, err
}

//GetPublicEntry returns a single complete PublicEntry with the specified ID
func (a ArticleSharingService) GetPublicEntry(entryID int) (PublicEntry, error) {
	entry, err := a.sharingRepo.FetchPublicEntry(entryID)
	if err != nil {
		return PublicEntry{}, err
	}
	article, err := a.sharingRepo.FetchArticle(entry.ArticleID)
	comments, err := a.sharingRepo.FetchComments(entryID, 0, 0)
	entry.ContainedArticle = &article
	entry.Comments = &comments
	return entry, err
}

//AddPublicEntry adds a single public entry into the database and returns it
func (a ArticleSharingService) AddPublicEntry(sourceRssID int, sharedBy, title, link string) (PublicEntry, error) {
	newArticle, err := a.addArticle(sourceRssID, title, link)
	if err != nil {
		return PublicEntry{}, err
	}
	newEntry := PublicEntry{
		ArticleID:  newArticle.ArticleID,
		SharedBy:   sharedBy,
		SharedDate: time.Now(),
	}
	newID, err := a.sharingRepo.AddPublicEntry(&newEntry)
	newEntry.EntryID = newID
	return newEntry, err
}

//ShareToUser adds a single recommendation and returns it
func (a ArticleSharingService) ShareToUser(senderUsername, receiverUsername, message string, articleID int) (Recommendation, error) {
	newRecommendation := Recommendation{
		SenderUsername:   senderUsername,
		ReceiverUsername: receiverUsername,
		Message:          message,
		ArticleID:        articleID,
	}
	err := a.sharingRepo.AddRecommendation(newRecommendation)
	return newRecommendation, err
}

//addArticle is a utility function that adds a new article instance to the db and returns it
func (a ArticleSharingService) addArticle(sourceRssID int, title, link string) (Article, error) {
	newArticle := Article{
		SourceRSS: sourceRssID,
		Title:     title,
		Link:      link,
	}
	newArticleID, err := a.sharingRepo.AddArticle(&newArticle)
	newArticle.ArticleID = newArticleID
	return newArticle, err
}
