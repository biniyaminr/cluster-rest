package sharing

import "time"

type SortOrder int

const (
	ASC SortOrder = iota
	DEC SortOrder = iota
)

//PublicEntry represents an article that has been shared to the public
type PublicEntry struct {
	EntryID          string
	ArticleID        string
	ContainedArticle *Article
	SharedBy         string
	SharedDate       time.Time
	LikesCount       int
	Comments         *[]Comment
}

//Article represents an article that has been saved in the database
type Article struct {
	ArticleID string
	SourceRSS string
	Title     string
	Link      string
}

//Comment represents a comment that has been posted to a public entry
type Comment struct {
	PostedBy    string
	CommentedOn time.Time
	Contents    string
}

//Recommendation represents a recommendation from one user to another
type Recommendation struct {
	SenderUsername   string
	ReceiverUsername string
	Article          *Article
}
