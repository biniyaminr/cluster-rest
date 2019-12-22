package sharing

import "time"

type SortOrder string
type SortBy string

const (
	ASC SortOrder = "ASC"
	DEC SortOrder = "DESC"
)

const (
	SHARED_DATE SortBy = "shared_date"
	LIKE_COUNT  SortBy = "likes_count"
)

//PublicEntry represents an article that has been shared to the public
type PublicEntry struct {
	EntryID          int
	ArticleID        int
	ContainedArticle *Article
	SharedBy         string
	SharedDate       time.Time
	LikesCount       int
	Comments         *[]Comment
}

//Article represents an article that has been saved in the database
type Article struct {
	ArticleID int
	SourceRSS int
	Title     string
	Link      string
}

//Comment represents a comment that has been posted to a public entry
type Comment struct {
	CommentID   int
	PostedBy    string
	CommentedOn time.Time
	Contents    string
}

//Recommendation represents a recommendation from one user to another
type Recommendation struct {
	SenderUsername   string
	ReceiverUsername string
	ArticleID        int
	Article          *Article
	Message          string
}
